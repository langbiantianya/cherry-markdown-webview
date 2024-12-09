export namespace file {
	
	export class File {
	    Name: string;
	    Path: string;
	    Bytes: number[];
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Path = source["Path"];
	        this.Bytes = source["Bytes"];
	    }
	}

}

