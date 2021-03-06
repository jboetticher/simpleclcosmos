import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeletePrice } from "./types/clprice/tx";
import { MsgCreatePrice } from "./types/clprice/tx";
import { MsgUpdateSentPrice } from "./types/clprice/tx";
import { MsgUpdatePrice } from "./types/clprice/tx";
import { MsgDeleteSentPrice } from "./types/clprice/tx";
import { MsgSendIbcPrice } from "./types/clprice/tx";
import { MsgCreateSentPrice } from "./types/clprice/tx";
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgDeletePrice: (data: MsgDeletePrice) => EncodeObject;
    msgCreatePrice: (data: MsgCreatePrice) => EncodeObject;
    msgUpdateSentPrice: (data: MsgUpdateSentPrice) => EncodeObject;
    msgUpdatePrice: (data: MsgUpdatePrice) => EncodeObject;
    msgDeleteSentPrice: (data: MsgDeleteSentPrice) => EncodeObject;
    msgSendIbcPrice: (data: MsgSendIbcPrice) => EncodeObject;
    msgCreateSentPrice: (data: MsgCreateSentPrice) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
