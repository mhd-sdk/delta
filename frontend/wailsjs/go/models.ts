export namespace persistence {
	
	export class Preferences {
	    theme: string;
	
	    static createFrom(source: any = {}) {
	        return new Preferences(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	    }
	}
	export class UserData {
	    username: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new UserData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.password = source["password"];
	    }
	}
	export class AppData {
	    user: UserData;
	    preferences: Preferences;
	
	    static createFrom(source: any = {}) {
	        return new AppData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.user = this.convertValues(source["user"], UserData);
	        this.preferences = this.convertValues(source["preferences"], Preferences);
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

export namespace rti {
	
	export class ResponseProductCodes {
	    template_id?: number;
	    user_msg?: string[];
	    rq_handler_rp_code?: string[];
	    rp_code?: string[];
	    exchange?: string;
	    symbol_name?: string;
	    product_code?: string;
	    timezone_time_of_interest?: string;
	    begin_time_of_interest_msm?: number;
	    end_time_of_interest_msm?: number;
	
	    static createFrom(source: any = {}) {
	        return new ResponseProductCodes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.template_id = source["template_id"];
	        this.user_msg = source["user_msg"];
	        this.rq_handler_rp_code = source["rq_handler_rp_code"];
	        this.rp_code = source["rp_code"];
	        this.exchange = source["exchange"];
	        this.symbol_name = source["symbol_name"];
	        this.product_code = source["product_code"];
	        this.timezone_time_of_interest = source["timezone_time_of_interest"];
	        this.begin_time_of_interest_msm = source["begin_time_of_interest_msm"];
	        this.end_time_of_interest_msm = source["end_time_of_interest_msm"];
	    }
	}

}

