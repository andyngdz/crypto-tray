export namespace config {
	
	export class Config {
	    provider_id: string;
	    api_keys: Record<string, string>;
	    refresh_seconds: number;
	    symbol: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.provider_id = source["provider_id"];
	        this.api_keys = source["api_keys"];
	        this.refresh_seconds = source["refresh_seconds"];
	        this.symbol = source["symbol"];
	    }
	}

}

export namespace main {
	
	export class ProviderInfo {
	    id: string;
	    name: string;
	    requiresApiKey: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ProviderInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.requiresApiKey = source["requiresApiKey"];
	    }
	}

}

