import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "darthlukan.starporttutorialapp.ibcdex";
export interface IbcdexPacketData {
    noData: NoData | undefined;
    /** this line is used by starport scaffolding # ibc/packet/proto/field */
    buyOrderPacket: BuyOrderPacketData | undefined;
}
export interface NoData {
}
/**
 * this line is used by starport scaffolding # ibc/packet/proto/message
 * BuyOrderPacketData defines a struct for the packet payload
 */
export interface BuyOrderPacketData {
    amountDenom: string;
    amount: number;
    priceDenom: string;
    price: number;
}
/** BuyOrderPacketAck defines a struct for the packet acknowledgment */
export interface BuyOrderPacketAck {
    remainingAmount: number;
    purchase: number;
}
export declare const IbcdexPacketData: {
    encode(message: IbcdexPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): IbcdexPacketData;
    fromJSON(object: any): IbcdexPacketData;
    toJSON(message: IbcdexPacketData): unknown;
    fromPartial(object: DeepPartial<IbcdexPacketData>): IbcdexPacketData;
};
export declare const NoData: {
    encode(_: NoData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NoData;
    fromJSON(_: any): NoData;
    toJSON(_: NoData): unknown;
    fromPartial(_: DeepPartial<NoData>): NoData;
};
export declare const BuyOrderPacketData: {
    encode(message: BuyOrderPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BuyOrderPacketData;
    fromJSON(object: any): BuyOrderPacketData;
    toJSON(message: BuyOrderPacketData): unknown;
    fromPartial(object: DeepPartial<BuyOrderPacketData>): BuyOrderPacketData;
};
export declare const BuyOrderPacketAck: {
    encode(message: BuyOrderPacketAck, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BuyOrderPacketAck;
    fromJSON(object: any): BuyOrderPacketAck;
    toJSON(message: BuyOrderPacketAck): unknown;
    fromPartial(object: DeepPartial<BuyOrderPacketAck>): BuyOrderPacketAck;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
