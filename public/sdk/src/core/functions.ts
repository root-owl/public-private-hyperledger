import FabricComm from '../core/query-invoke';
import { ResponseModel } from '../core/model';
const fabricComm = new FabricComm();
import * as fs from 'fs';
import * as path from 'path';

const connectionProfilePath = path.resolve(__dirname, '..', '..', 'connection.json')
const cpJSON = fs.readFileSync(connectionProfilePath, 'utf-8');
const connectionProfile = JSON.parse(cpJSON);
const channelName = "public.secc";
const chaincodeName = "secc";

enum ChaincodeFunctions {
   Login = "UserLogin"
}

export default class ChaincodeFunc {
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
}