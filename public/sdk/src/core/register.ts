import FabricCAServices = require('fabric-ca-client');
import { BaseClient } from './base';

interface ProtoBufObject {
    toBuffer(): Buffer;
}

declare class Remote {
    constructor(url: string, opts?: Client.ConnectionOpts);
    public getName(): string;
    public setName(name: string): void;
    public getUrl(): string;
}

declare class Client extends BaseClient {
    public static loadFromConfig(config: any): Client;

    constructor();
    public loadFromConfig(config: any): void;
    public setTlsClientCertAndKey(clientCert: string, clientKey: string): void;
    public addTlsClientCertAndKey(opts: any): void;
    public isDevMode(): boolean;
    public setDevMode(mode: boolean): void;
    public newChannel(name: string): Client.Channel;
    public getChannel(name?: string, throwError?: boolean): Client.Channel;
    public newPeer(url: string, opts?: Client.ConnectionOpts): Client.Peer;
    public getPeer(name: string): Client.Peer;
    public getPeersForOrg(mspid?: string): Client.Peer[];
    public newOrderer(url: string, opts?: Client.ConnectionOpts): Client.Orderer;
    public getOrderer(name: string): Client.Orderer;
    public getPeersForOrgOnChannel(channelNames: string | string[]): Client.ChannelPeer[];
    public getCertificateAuthority(): FabricCAServices;
    public getClientConfig(): any;
    public getMspid(): string;
    public newTransactionID(admin?: boolean): Client.TransactionId;
    public extractChannelConfig(configEnvelope: Buffer): Buffer;
    public signChannelConfig(config: Buffer): Client.ConfigSignature;
    public createChannel(request: Client.ChannelRequest): Promise<Client.BroadcastResponse>;
    public updateChannel(request: Client.ChannelRequest): Promise<Client.BroadcastResponse>;
    public queryPeers(request: Client.PeerQueryRequest): Promise<Client.PeerQueryResponse>;
    public queryChannels(peer: Client.Peer | string, useAdmin?: boolean): Promise<Client.ChannelQueryResponse>;
    public queryInstalledChaincodes(peer: Client.Peer | string, useAdmin?: boolean): Promise<Client.ChaincodeQueryResponse>;
    public installChaincode(request: Client.ChaincodeInstallRequest, timeout?: number): Promise<Client.ProposalResponseObject>;
    public initCredentialStores(): Promise<boolean>;
    public setStateStore(store: Client.IKeyValueStore): void;
    public setAdminSigningIdentity(privateKey: string, certificate: string, mspid: string): void;
    public saveUserToStateStore(): Promise<Client.User>;
    public setUserContext(user: Client.User | Client.UserContext, skipPersistence?: boolean): Promise<Client.User>;
    public getUserContext(name: string, checkPersistence?: boolean): Promise<Client.User> | Client.User;
    public loadUserFromStateStore(name: string): Promise<Client.User>;
    public getStateStore(): Client.IKeyValueStore;
    public createUser(opts: Client.UserOpts): Promise<Client.User>;

    public getTargetPeers(requestTargets: string | string[] | Client.Peer | Client.Peer[]): Client.Peer[];
    public getTargetOrderer(requestOrderer?: string | Client.Orderer, channelOrderers?: Client.Orderer[], channelName?: string): Client.Orderer;
    public getClientCertHash(create: boolean): Buffer;
}

export = Client;

declare namespace Client { // tslint:disable-line:no-namespace
    export enum Status {
        UNKNOWN = 0,
        SUCCESS = 200,
        BAD_REQUEST = 400,
        FORBIDDEN = 403,
        NOT_FOUND = 404,
        REQUEST_ENTITY_TOO_LARGE = 413,
        INTERNAL_SERVER_ERROR = 500,
        SERVICE_UNAVAILABLE = 503,
    }

    export type ChaincodeType = 'golang' | 'car' | 'java' | 'node';
    export interface ICryptoKey {
        getSKI(): string;
        isSymmetric(): boolean;
        isPrivate(): boolean;
        getPublicKey(): ICryptoKey;
        toBytes(): string;
    }

    export interface ICryptoKeyStore {
        getKey(ski: string): Promise<string>;
        putKey(key: ICryptoKey): Promise<ICryptoKey>;
    }

    export interface ICryptoSuite {
        decrypt(key: ICryptoKey, cipherText: Buffer, opts: any): Buffer;
        deriveKey(key: ICryptoKey, opts?: KeyOpts): ICryptoKey;
        encrypt(key: ICryptoKey, plainText: Buffer, opts: any): Buffer;
        getKey(ski: string): Promise<ICryptoKey>;
        generateKey(opts?: KeyOpts): Promise<ICryptoKey>;
        hash(msg: string, opts: any): string;
        importKey(pem: string, opts?: KeyOpts): ICryptoKey | Promise<ICryptoKey>;
        setCryptoKeyStore(cryptoKeyStore: ICryptoKeyStore): void;
        sign(key: ICryptoKey, digest: Buffer): Buffer;
        verify(key: ICryptoKey, signature: Buffer, digest: Buffer): boolean;
    }

    export interface CryptoSetting {
        algorithm: string;
        hash: string;
        keysize: number;
        software: boolean;
    }

    export interface UserConfig {
        affiliation?: string;
        enrollmentID: string;
        name: string;
        roles?: string[];
    }

    export interface ConnectionOpts {
        pem?: string;
        clientKey?: string;
        clientCert?: string;
        'request-timeout'?: string;
        'ssl-target-name-override'?: string;
        [propName: string]: any;
    }

    export class User {
        public static isInstance(object: any): boolean;

        constructor(cfg: string | UserConfig);
        public getName(): string;
        public getRoles(): string[];
        public setRoles(roles: string[]): void;
        public getAffiliation(): string;
        public setAffiliation(affiliation: string): void;
        public getIdentity(): IIdentity;
        public getSigningIdentity(): ISigningIdentity;
        public getCryptoSuite(): ICryptoSuite;
        public setCryptoSuite(suite: ICryptoSuite): void;
        public setEnrollment(privateKey: ICryptoKey, certificate: string, mspId: string): Promise<void>;
        public isEnrolled(): boolean;
        public fromString(): Promise<User>;
    }