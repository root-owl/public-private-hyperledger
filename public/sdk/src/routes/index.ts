import * as express from 'express';

export const indexRouter = express.Router();

indexRouter.get('/', function (req: express.Request, res: express.Response, next: express.NextFunction) {
    res.render('index', { title: 'Expresss' });
})