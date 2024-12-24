export namespace config {
	
	export class AliyunOSS {
	    accessKeyID: string;
	    accessKeySecret: string;
	    region: string;
	    bucketName: string;
	
	    static createFrom(source: any = {}) {
	        return new AliyunOSS(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accessKeyID = source["accessKeyID"];
	        this.accessKeySecret = source["accessKeySecret"];
	        this.region = source["region"];
	        this.bucketName = source["bucketName"];
	    }
	}
	export class TencentCloudCOS {
	    secretID: string;
	    secretKey: string;
	    bucketURL: string;
	
	    static createFrom(source: any = {}) {
	        return new TencentCloudCOS(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.secretID = source["secretID"];
	        this.secretKey = source["secretKey"];
	        this.bucketURL = source["bucketURL"];
	    }
	}
	export class PicBed {
	    basePath: string;
	    activated: string;
	    oss?: AliyunOSS;
	    cos?: TencentCloudCOS;
	
	    static createFrom(source: any = {}) {
	        return new PicBed(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.basePath = source["basePath"];
	        this.activated = source["activated"];
	        this.oss = this.convertValues(source["oss"], AliyunOSS);
	        this.cos = this.convertValues(source["cos"], TencentCloudCOS);
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

