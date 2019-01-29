import {logger} from './logger';

export default class Utils{
    private static instance: Utils

    public static getInstance(): Utils{
        if (this.instance == null){
            this.instance = new Utils();
        }

        return this.instance;
    }

    public log(val: any){
        logger.debug(val)
    }
}