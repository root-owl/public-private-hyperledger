import * as express from 'express';
import ChaincodeFunc from '../core/functions';

export const authRouter = express.Router();

authRouter.post('/login', async function (req: express.Request, res: express.Response, next: express.NextFunction) {
  let response = await ChaincodeFunc.loginUser(req.body.email, req.body.password, req.body.role)
  let code = 200;
  if (response.code != undefined) {
    code = response.code;
  }

  res.status(code);
  res.send(response)
})

authRouter.post('/register', async function (req: express.Request, res: express.Response, next: express.NextFunction) {
  let response = await ChaincodeFunc.registerUser(req.body.first_name,
    req.body.last_name,
    req.body.public_key,
    req.body.public_address,
    req.body.description,
    req.body.email,
    req.body.role)

  let code = 200;
  if (response.code != undefined) {
    code = response.code
  }

  res.status(code)
  res.send(response)
})
