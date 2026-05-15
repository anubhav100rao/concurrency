# Locking on Ticket Booking system, drivers to orders matching such that only one driver can accept the order

## The Database atomic update
```sql
UPDATE orders
SET status = 'ACCEPTED', driver_id = 1
WHERE id = 1 AND status = 'PENDING';
```

## MVCC (multi version concurrency control)
```sql
Update seats SET status = 'BOOKED', version=6 
WHERE seat_number = 4 AND version = 5;
```

## Distributed lock, when `state` isn't just a single DB row, multiple DBs involved
Imagine an e-commerce checkout process:

- Check if the item is in stock.
- Call Stripe's API to charge the user's credit card.
- Update the database to say the item is sold.

If two users click "Buy" at the exact same time, and you rely on a database lock at step 3, both users will pass step 1, both will be charged on Stripe in step 2, and then step 3 will fail for one of them. You've double-charged a customer.

To prevent this we need to lock entire process:
`acquire lock -> check stock -> call stripe -> update DB -> release lock`

```javascript
async function processCheckout(orderId) {
    // 1. Try to grab the lock in Redis (Ask for 5 seconds of TTL)
    const lock = await redis.set(`lock_${orderId}`, "Server_A", { NX: true, PX: 5000 });
    
    if (!lock) {
        return "Someone else is processing this order!"; 
    }

    try {
        // --- CRITICAL SECTION BEGINS ---
        await db.query("UPDATE orders SET status = 'PROCESSING' WHERE id = ?", [orderId]);
        
        const stripeResult = await stripe.chargeCard(50.00); 
        
        await db.query("UPDATE payments SET status = 'PAID' WHERE id = ?", [stripeResult.id]);
        // --- CRITICAL SECTION ENDS ---
    } finally {
        // 3. Delete the lock manually when done
        await redis.del(`lock_${orderId}`);
    }
}
```

```javascript
async function processCheckout(orderId) {
    // 1. Zookeeper client establishes a session (Background heartbeats start automatically)
    // 2. Create an "Ephemeral" (temporary) lock node
    const lockNode = await zk.createEphemeral(`/locks/order_${orderId}`);
    
    if (!lockNode) {
        return "Someone else is processing this order!"; 
    }

    try {
        // --- CRITICAL SECTION BEGINS ---
        await db.query("UPDATE orders SET status = 'PROCESSING' WHERE id = ?", [orderId]);
        
        const stripeResult = await stripe.chargeCard(50.00); 
        
        await db.query("UPDATE payments SET status = 'PAID' WHERE id = ?", [stripeResult.id]);
        // --- CRITICAL SECTION ENDS ---
    } finally {
        // 3. Delete the node manually when done
        await zk.delete(`/locks/order_${orderId}`);
    }
}
```