import { stat } from "fs";

export class ResponseModel {
    message: string
    status?: number

    constructor(message?: string, status?: number) {
        this.message = message;
        this.status = status;
    }
}