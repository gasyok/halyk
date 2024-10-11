# Test assessment

## Entities

User, Lot, Bid, Auction

Connections between entities:

User : left_right_arrow: Lot:

- One user (seller) can create multiple lots
- Lot belongs to only one seller

User : left_right_arrow: Bid:

- One user (buyer) can create multiple bids
- Bid belongs to only one user

Lot: left_right_arrow: Auction:

- Each lot is connected to one auction
- Auction is held only for one lot

Auction: left_right_arrow: Bid:

- Auction has multiple bids
- Bid refers to only one auction

Auction: left_right_arrow: Winner:

- After the end of auction the winner is determined

А так идите нахуй
