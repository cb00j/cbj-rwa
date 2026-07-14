sequenceDiagram
autonumber
participant U as "用户 (User)"
participant G as "CBJGateway 合约 (On-chain)"
participant B as "Backend & Broker (Off-chain)"

    %% ---------------- 存入流程 ----------------
    Note over U,G: 📥 存入流程 (Deposit)

    U->>G: 1. 存入 USDC
    G-->>U: emit PendingDeposit & mint PendingCbjUSDC & (状态: pending)

    G-->>B: 2. 发送 PendingDeposit 事件
    Note over B: Broker 入账处理 (Off-chain)
    B-->>G: burn PendingCbjUSDC & mint cbjUSDC & (状态: active)
    G-->>U: 返回 cbjUSDC

    %% ---------------- 取出流程 ----------------
    Note over U,G: 📤 取出流程 (Withdraw)

    U->>G: 1. 存入 cbjUSDC
    G-->>U: burn cbjUSDC & emit PendingWithdraw & mint PendingUSDC & (状态: pending)

    G-->>B: 2. 发送 PendingWithdraw 事件
    Note over B: Broker 赎回 USDC (Off-chain)
    B-->>G: burn PendingUSDC & USDC 转回 CBJGateway
    G-->>U: 返回 USDC (状态: redeemed)
