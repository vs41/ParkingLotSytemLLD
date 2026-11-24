# Parking Lot LLD (in-memory)

## Run

```bash
go run .
go test ./...
```

## Design notes

- `ParkingService` coordinates entry and exit workflows.
- `PaymentService` is a struct that owns a registry of `PaymentStrategy` implementations. The customer chooses `UPI`, `CARD`, or `CASH` for each exit/payment request.
- `FeeCalculatorService` owns pricing strategies. `ParkingLot.PricingType` selects `HOURLY` or `DAILY` pricing for that lot.
- `SpotAllocationService` is the only module responsible for finding, allocating, and releasing spots.
- `TicketService` owns ticket creation and lifecycle updates.
- **State Pattern:** `ParkingSpot` delegates allocation/removal behavior to its current state (`AVAILABLE`, `OCCUPIED`, or `OUT_OF_SERVICE`). Invalid operations are rejected by the active state.
- **No full State Pattern for Ticket/Payment:** their current lifecycle is simple, so they use typed status values and validated transition methods (`Ticket.Complete`, `Payment.MarkSuccessful`, `Payment.MarkFailed`). Add dedicated state objects only if those states gain substantially different behavior.
- `main.go` creates one in-memory database and injects it into both repositories.

This is a learning/demo implementation. Payments are simulated as successful; there is no real payment gateway. The in-memory spot allocation is not safe for concurrent multi-process production traffic. A production implementation should use a database transaction/conditional update for spot reservation, persist pricing snapshots on tickets, and store currency in integer minor units rather than `float64`.


![alt text](<mermaid-diagram (1).png>)