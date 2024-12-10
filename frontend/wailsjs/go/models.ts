export namespace file {
	
	export class File {
	    DisplayName: string;
	    Pattern: string;
	    Name: string;
	    Path: string;
	    Bytes: number[];
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DisplayName = source["DisplayName"];
	        this.Pattern = source["Pattern"];
	        this.Name = source["Name"];
	        this.Path = source["Path"];
	        this.Bytes = source["Bytes"];
	    }
	}

}

