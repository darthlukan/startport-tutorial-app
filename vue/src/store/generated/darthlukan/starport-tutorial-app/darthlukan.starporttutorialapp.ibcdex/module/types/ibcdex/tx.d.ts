import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "darthlukan.starporttutorialapp.ibcdex";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgSendBuyOrder {
    sender: string;
    port: string;
    channelID: string;
    timeoutTimestamp: number;
    amountDenom: string;
    amount: number;
    priceDenom: string;
    price: number;
}
export interface MsgSendBuyOrderResponse {
}
export declare const MsgSendBuyOrder: {
    encode(message: MsgSendBuyOrder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendBuyOrder;
    fromJSON(object: any): MsgSendBuyOrder;
    toJSON(message: MsgSendBuyOrder): unknown;
    fromPartial(object: DeepPartial<MsgSendBuyOrder>): MsgSendBuyOrder;
};
export declare const MsgSendBuyOrderResponse: {
    encode(_: MsgSendBuyOrderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendBuyOrderResponse;
    fromJSON(_: any): MsgSendBuyOrderResponse;
    toJSON(_: MsgSendBuyOrderResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendBuyOrderResponse>): MsgSendBuyOrderResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SendBuyOrder(request: MsgSendBuyOrder): Promise<MsgSendBuyOrderResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    SendBuyOrder(request: MsgSendBuyOrder): Promise<MsgSendBuyOrderResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
