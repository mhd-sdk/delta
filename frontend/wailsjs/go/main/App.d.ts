// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {alpaca} from '../models';
import {persistence} from '../models';
import {main} from '../models';
import {marketdata} from '../models';

export function GetAccount():Promise<alpaca.Account>;

export function GetAppData():Promise<persistence.AppData>;

export function GetAssets():Promise<Array<alpaca.Asset>>;

export function GetCandlesticks(arg1:main.GetCandlesticksConfig):Promise<Array<marketdata.Bar>>;

export function Logout():Promise<void>;

export function ResetPreferences():Promise<void>;

export function SaveAppData(arg1:persistence.AppData):Promise<void>;

export function TestCredentials(arg1:string,arg2:string):Promise<boolean>;
