import * as fs from 'fs';
import * as path from 'path';
import * as uuid from 'uuid/v1';
import FabricComm from '../core/query-invoke';
import { ResponseModel } from '../core/model';
import CAClient from '../core/fabric-ca-client';
const fabricComm = new FabricComm();


const connectionProfilePath = path.resolve(__dirname, '..', '..', 'connection.json')
const cpJSON = fs.readFileSync(connectionProfilePath, 'utf-8');
const connectionProfile = JSON.parse(cpJSON);
const channelName = "public.secc";
const chaincodeName = "secc";

const fabricCA = new CAClient();

enum ChaincodeFunctions {
   Login = "UserLogin",
   Register = "UserRegister"
}

export default class ChaincodeFunc {
   //Login User
   static async loginUser(email: string, password: string, role: number): Promise<any> {
      try {
         let loginResponse = await fabricComm.query(channelName,
            'admin',
            ChaincodeFunctions.Login,
            chaincodeName,
            { email: email, password: password, role: role })

         return loginResponse
      } catch (error) {
         let loginResponse = new ResponseModel(error.message, 500);
         return loginResponse;
      }
   }

   //Register user
   static async registerUser(firstName: string, lastName: string,
      publicKey: string, publicAddress: string, description: string,
      email: string, role: number): Promise<any> {

      try {
         const id = uuid();
         const data = {
            member_id: id,
            first_name: firstName,
            last_name: lastName,
            public_key: publicKey,
            public_address: "",
            description: "",
            email_address: email,
            password: "abc123",
            role: role
         };
         await fabricCA.enrollUser();
         await fabricCA.registerUser(email, 'admin@123');
         let signupResponse = await fabricComm.invoke(channelName, email, ChaincodeFunctions.Register,
            data, chaincodeName);
         return signupResponse;
      } catch (error) {
         let loginResponse = new ResponseModel(error.message, 500);
         return loginResponse;
      }
   }
}