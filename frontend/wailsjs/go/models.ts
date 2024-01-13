export namespace models {
	
	export class Login {
	    username: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new Login(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.password = source["password"];
	    }
	}
	export class Register {
	    username: string;
	    password: string;
	    password2: string;
	    email: string;
	
	    static createFrom(source: any = {}) {
	        return new Register(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.password = source["password"];
	        this.password2 = source["password2"];
	        this.email = source["email"];
	    }
	}
	export class User {
	    id: number[];
	    first_name: string;
	    username: string;
	    email: string;
	    jwt: string;
	    password: string;
	    // Go type: time
	    modified_at: any;
	    // Go type: time
	    created_at: any;
	    // Go type: gorm
	    deleted_at: any;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.first_name = source["first_name"];
	        this.username = source["username"];
	        this.email = source["email"];
	        this.jwt = source["jwt"];
	        this.password = source["password"];
	        this.modified_at = this.convertValues(source["modified_at"], null);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.deleted_at = this.convertValues(source["deleted_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

