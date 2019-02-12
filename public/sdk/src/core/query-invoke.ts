
import * as fs from 'fs';
import * as path from 'path';
import { ResponseModel } from './model';
import { FileSystemWallet, Gateway } from 'fabric-network';
import CAClient from './fabric-ca-client';

const connectionProfilePath = path.resolve(__dirname, '..', '..', 'connection.json')
const cpJSON = fs.readFileSync(connectionProfilePath, 'utf-8');
const connectionProfile = JSON.parse(cpJSON);

const walletPath = path.resolve(__dirname, '..', '..', 'wallet')

const wallet = new FileSystemWallet(walletPath);

export default class FabricComm {

    /**
     * This function will invoke chaincode function
     * @param channel - name of the channel
     * @param username - username
     * @param functionName - chaincode function name
     * @param data - request payload
     * @param chaincode - chaincode name
     */
    public async invoke(channel: string,
        username: string,
        functionName: string,
        data: any,
        chaincode: string): Promise<any> {
        console.log(connectionProfilePath)
        try {
            const userExistenceErr = await this.checkUserExistence(username);

            if (userExistenceErr != null) {
                return new ResponseModel(userExistenceErr.message, 500);
            }

            // Create a new gateway for connecting to our peer node.
            const gateway = new Gateway();
            await gateway.connect(connectionProfile, { wallet, identity: username, discovery: { asLocalhost: false, enabled: false } });
            // // Get the network (channel) our contract is deployed to.
            const network = await gateway.getNetwork(channel);
            // // Get the contract from the network.
            const contract = network.getContract(chaincode);
            const respBuffer = await contract.submitTransaction(functionName, JSON.stringify(data));
            gateway.disconnect();
            return JSON.parse(respBuffer.toString());
            // return "ok"
        } catch (error) {
            console.log(error)
            let response = new ResponseModel(error.message, 500);
            return response;
        }
    }

    /**
     * This function will query blockchain
     * @param channel - Channel name
     * @param username - username of client
     * @param functionName - name of function to invoke
     * @param chaincode - chaincode name
     * @param queryData - request payload if any
     */
    public async query(channel: string,
        username: string,
        functionName: string,
        chaincode: string,
        queryData?: any
    ): Promise<any> {
        try {
            const userExistenceErr = await this.checkUserExistence(username);

            if (userExistenceErr != null) {
                return new ResponseModel(userExistenceErr.message, 401);
            }

            let caClient = new CAClient();
            await caClient.enrollUser()

            // Create a new gateway for connecting to our peer node.
            const gateway = new Gateway();
            await gateway.connect(connectionProfile, { wallet, identity: username, discovery: { enabled: true } });
            // Get the network (channel) our contract is deployed to.
            const network = await gateway.getNetwork(channel);
            // Get the contract from the network.
            const contract = network.getContract(chaincode);
            let respBuffer: Buffer
            if (queryData != null && queryData != undefined) {
                respBuffer = await contract.evaluateTransaction(functionName, JSON.stringify(queryData));
            } else {
                respBuffer = await contract.evaluateTransaction(functionName);
            }
            return JSON.parse(respBuffer.toString())
        } catch (error) {
            console.log(error)
            let response = new ResponseModel(error.message, 500);
            return response;
        }
    }

    /**
     * This function will check whether user certs exist in wallet or not
     * @param username username of user
     */
    private checkUserExistence(username: string): Promise<Error> {
        try {
            const userExist = wallet.exists(username);

            if (!userExist) {
                throw new Error("Please enroll " + username + " first")
            }

        } catch (error) {
            return error
        }

    }
}

