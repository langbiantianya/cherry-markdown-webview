export namespace config {
	
	export class Aliyun {
	    accessKeyID: string;
	    accessKeySecret: string;
	    endpoint: string;
	    bucketName: string;
	
	    static createFrom(source: any = {}) {
	        return new Aliyun(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accessKeyID = source["accessKeyID"];
	        this.accessKeySecret = source["accessKeySecret"];
	        this.endpoint = source["endpoint"];
	        this.bucketName = source["bucketName"];
	    }
	}
	export class TencentCloud {
	    secretID: string;
	    secretKey: string;
	    region: string;
	    bucketName: string;
	
	    static createFrom(source: any = {}) {
	        return new TencentCloud(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.secretID = source["secretID"];
	        this.secretKey = source["secretKey"];
	        this.region = source["region"];
	        this.bucketName = source["bucketName"];
	    }
	}
	export class OSS {
	    basePath: string;
	    aliyun?: Aliyun;
	    tencentCloud?: TencentCloud;
	
	    static createFrom(source: any = {}) {
	        return new OSS(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.basePath = source["basePath"];
	        this.aliyun = this.convertValues(source["aliyun"], Aliyun);
	        this.tencentCloud = this.convertValues(source["tencentCloud"], TencentCloud);
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

export namespace file {
	
	export class File {
	    Mime: string;
	    DisplayName: string;
	    Pattern: string;
	    Name: string;
	    Path: string;
	    Bytes: number[];
	    StrData: string;
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mime = source["Mime"];
	        this.DisplayName = source["DisplayName"];
	        this.Pattern = source["Pattern"];
	        this.Name = source["Name"];
	        this.Path = source["Path"];
	        this.Bytes = source["Bytes"];
	        this.StrData = source["StrData"];
	    }
	}

}

