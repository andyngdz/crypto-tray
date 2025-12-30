export namespace config {
	
	export class Config {
	    provider_id: string;
	    api_keys: Record<string, string>;
	    refresh_seconds: number;
	    symbols: string[];
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.provider_id = source["provider_id"];
	        this.api_keys = source["api_keys"];
	        this.refresh_seconds = source["refresh_seconds"];
	        this.symbols = source["symbols"];
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

export namespace providers {
	
	export class SymbolInfo {
	    id: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new SymbolInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}

}

