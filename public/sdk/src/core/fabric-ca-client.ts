import * as FabricCAServices from 'fabric-ca-client';
import { FileSystemWallet, X509WalletMixin, Gateway } from 'fabric-network';
import * as fs from 'fs';
import * as path from 'path';

const connectionProfilePath = path.resolve(__dirname, '..', '..', 'connection.json')
const cpJSON = fs.readFileSync(connectionProfilePath, 'utf-8');
const connectionProfile = JSON.parse(cpJSON);

const walletPath = path.resolve(__dirname, '..', '..', 'wallet')

const caURL = connectionProfile.certificateAuthorities['ca.public.secc.com'].url;
const ca = new FabricCAServices(caURL);

const wallet = new FileSystemWallet(walletPath);


export default class CAClient {

    /**
     * Enroll user
     * @param username username of user
     * @param password secret 
     */
    public async enrollUser(username: string = 'admin', password: string = 'adminpw'): Promise<Error> {
        try {

            const exist = await wallet.exists('admin');
            if (exist) {
                // Utils.getInstance().log("Admin already enrolled");
                return null;
            }

            const enrollment = await ca.enroll({ enrollmentID: username, enrollmentSecret: password })
            const identity = X509WalletMixin.createIdentity('SECCPeerOrgMSP', enrollment.certificate, enrollment.key.toBytes());
            wallet.import(username, identity);
            // Utils.getInstance().log("Enrollment of admin done.")
            return null
        } catch (error) {
            console.log(error)
            return error
        }
    }

    /**
     * Register & enroll user with CA
     * @param username - username
     * @param secret - secret key
     */
    public async registerUser(username: string, secret: string): Promise<Error> {
        try {

            const adminExist = await wallet.exists('admin')
            if (!adminExist) {
                return new Error("Please enroll admin first")
            }

            const exists = await wallet.exists(username)

            if (exists) {
                return new Error("User already exist")
            }

            // Create a new gateway for connecting to our peer node.
            const gateway = new Gateway();
            await gateway.connect(connectionProfile, { wallet, identity: 'admin', discovery: { enabled: false } });

            // Get the CA client object from the gateway for interacting with the CA.
            const ca = gateway.getClient().getCertificateAuthority();
            const adminIdentity = gateway.getCurrentIdentity();


            await ca.register({ enrollmentID: username, enrollmentSecret: secret, role: 'client', affiliation: 'org1.department1' }, adminIdentity)
            const enrollment = await ca.enroll({ enrollmentID: username, enrollmentSecret: secret })
            const identity = X509WalletMixin.createIdentity('SECCPeerOrgMSP', enrollment.certificate, enrollment.key.toBytes());
            await wallet.import(username, identity);
        } catch (error) {
            console.log(error)
            return error
        }
    }
}

