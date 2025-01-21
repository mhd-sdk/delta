export namespace alpaca {
	
	export class Account {
	    id: string;
	    account_number: string;
	    status: string;
	    crypto_status: string;
	    currency: string;
	    // Go type: decimal
	    buying_power: any;
	    // Go type: decimal
	    regt_buying_power: any;
	    // Go type: decimal
	    daytrading_buying_power: any;
	    // Go type: decimal
	    effective_buying_power: any;
	    // Go type: decimal
	    non_marginable_buying_power: any;
	    // Go type: decimal
	    bod_dtbp: any;
	    // Go type: decimal
	    cash: any;
	    // Go type: decimal
	    accrued_fees: any;
	    // Go type: decimal
	    portfolio_value: any;
	    pattern_day_trader: boolean;
	    trading_blocked: boolean;
	    transfers_blocked: boolean;
	    account_blocked: boolean;
	    shorting_enabled: boolean;
	    trade_suspended_by_user: boolean;
	    // Go type: time
	    created_at: any;
	    // Go type: decimal
	    multiplier: any;
	    // Go type: decimal
	    equity: any;
	    // Go type: decimal
	    last_equity: any;
	    // Go type: decimal
	    long_market_value: any;
	    // Go type: decimal
	    short_market_value: any;
	    // Go type: decimal
	    position_market_value: any;
	    // Go type: decimal
	    initial_margin: any;
	    // Go type: decimal
	    maintenance_margin: any;
	    // Go type: decimal
	    last_maintenance_margin: any;
	    // Go type: decimal
	    sma: any;
	    daytrade_count: number;
	    crypto_tier: number;
	
	    static createFrom(source: any = {}) {
	        return new Account(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.account_number = source["account_number"];
	        this.status = source["status"];
	        this.crypto_status = source["crypto_status"];
	        this.currency = source["currency"];
	        this.buying_power = this.convertValues(source["buying_power"], null);
	        this.regt_buying_power = this.convertValues(source["regt_buying_power"], null);
	        this.daytrading_buying_power = this.convertValues(source["daytrading_buying_power"], null);
	        this.effective_buying_power = this.convertValues(source["effective_buying_power"], null);
	        this.non_marginable_buying_power = this.convertValues(source["non_marginable_buying_power"], null);
	        this.bod_dtbp = this.convertValues(source["bod_dtbp"], null);
	        this.cash = this.convertValues(source["cash"], null);
	        this.accrued_fees = this.convertValues(source["accrued_fees"], null);
	        this.portfolio_value = this.convertValues(source["portfolio_value"], null);
	        this.pattern_day_trader = source["pattern_day_trader"];
	        this.trading_blocked = source["trading_blocked"];
	        this.transfers_blocked = source["transfers_blocked"];
	        this.account_blocked = source["account_blocked"];
	        this.shorting_enabled = source["shorting_enabled"];
	        this.trade_suspended_by_user = source["trade_suspended_by_user"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.multiplier = this.convertValues(source["multiplier"], null);
	        this.equity = this.convertValues(source["equity"], null);
	        this.last_equity = this.convertValues(source["last_equity"], null);
	        this.long_market_value = this.convertValues(source["long_market_value"], null);
	        this.short_market_value = this.convertValues(source["short_market_value"], null);
	        this.position_market_value = this.convertValues(source["position_market_value"], null);
	        this.initial_margin = this.convertValues(source["initial_margin"], null);
	        this.maintenance_margin = this.convertValues(source["maintenance_margin"], null);
	        this.last_maintenance_margin = this.convertValues(source["last_maintenance_margin"], null);
	        this.sma = this.convertValues(source["sma"], null);
	        this.daytrade_count = source["daytrade_count"];
	        this.crypto_tier = source["crypto_tier"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Asset {
	    id: string;
	    class: string;
	    exchange: string;
	    symbol: string;
	    name: string;
	    status: string;
	    tradable: boolean;
	    marginable: boolean;
	    maintenance_margin_requirement: number;
	    shortable: boolean;
	    easy_to_borrow: boolean;
	    fractionable: boolean;
	    attributes: string[];
	
	    static createFrom(source: any = {}) {
	        return new Asset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.class = source["class"];
	        this.exchange = source["exchange"];
	        this.symbol = source["symbol"];
	        this.name = source["name"];
	        this.status = source["status"];
	        this.tradable = source["tradable"];
	        this.marginable = source["marginable"];
	        this.maintenance_margin_requirement = source["maintenance_margin_requirement"];
	        this.shortable = source["shortable"];
	        this.easy_to_borrow = source["easy_to_borrow"];
	        this.fractionable = source["fractionable"];
	        this.attributes = source["attributes"];
	    }
	}

}

export namespace app {
	
	export class TimeFrame {
	    N: number;
	    Unit: string;
	
	    static createFrom(source: any = {}) {
	        return new TimeFrame(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.N = source["N"];
	        this.Unit = source["Unit"];
	    }
	}
	export class GetCandlesticksConfig {
	    Ticker: string;
	    // Go type: time
	    Start: any;
	    // Go type: time
	    End: any;
	    timeframe: TimeFrame;
	
	    static createFrom(source: any = {}) {
	        return new GetCandlesticksConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Ticker = source["Ticker"];
	        this.Start = this.convertValues(source["Start"], null);
	        this.End = this.convertValues(source["End"], null);
	        this.timeframe = this.convertValues(source["timeframe"], TimeFrame);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace marketdata {
	
	export class Bar {
	    // Go type: time
	    t: any;
	    o: number;
	    h: number;
	    l: number;
	    c: number;
	    v: number;
	    n: number;
	    vw: number;
	
	    static createFrom(source: any = {}) {
	        return new Bar(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.t = this.convertValues(source["t"], null);
	        this.o = source["o"];
	        this.h = source["h"];
	        this.l = source["l"];
	        this.c = source["c"];
	        this.v = source["v"];
	        this.n = source["n"];
	        this.vw = source["vw"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace options {
	
	export class SecondInstanceData {
	    Args: string[];
	    WorkingDirectory: string;
	
	    static createFrom(source: any = {}) {
	        return new SecondInstanceData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Args = source["Args"];
	        this.WorkingDirectory = source["WorkingDirectory"];
	    }
	}

}

export namespace persistence {
	
	export class Tile {
	    ID: string;
	    Type: string;
	    X: number;
	    Y: number;
	    W: number;
	    H: number;
	
	    static createFrom(source: any = {}) {
	        return new Tile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Type = source["Type"];
	        this.X = source["X"];
	        this.Y = source["Y"];
	        this.W = source["W"];
	        this.H = source["H"];
	    }
	}
	export class Workspace {
	    name: string;
	    layout: Tile[];
	
	    static createFrom(source: any = {}) {
	        return new Workspace(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.layout = this.convertValues(source["layout"], Tile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GeneralPreferences {
	    theme: string;
	    language: string;
	
	    static createFrom(source: any = {}) {
	        return new GeneralPreferences(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	        this.language = source["language"];
	    }
	}
	export class Preferences {
	    generalPreferences: GeneralPreferences;
	
	    static createFrom(source: any = {}) {
	        return new Preferences(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.generalPreferences = this.convertValues(source["generalPreferences"], GeneralPreferences);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Keys {
	    apiKey: string;
	    secretKey: string;
	
	    static createFrom(source: any = {}) {
	        return new Keys(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.apiKey = source["apiKey"];
	        this.secretKey = source["secretKey"];
	    }
	}
	export class AppData {
	    keys: Keys;
	    preferences: Preferences;
	    workspaces: Workspace[];
	    favoriteTickers: string[];
	
	    static createFrom(source: any = {}) {
	        return new AppData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.keys = this.convertValues(source["keys"], Keys);
	        this.preferences = this.convertValues(source["preferences"], Preferences);
	        this.workspaces = this.convertValues(source["workspaces"], Workspace);
	        this.favoriteTickers = source["favoriteTickers"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	

}

