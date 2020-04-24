# PostmanToAPIInput
Convert Postman JSON file to Swift codes APIInput file.  
*This is my first experience coding with Go. So bear with the code.*

## How to use
The easiest one is using the build: 
    
**\# ./postmantoswift-executable**   
    
Postman file path: /path/to/postman/file.json  
Path to save the file: /path/to/save  
Base struct name (default APIInputType):  
*File saved to: /path/to/save/APIInput.swift*

The second one is execute the from the source code directly. You may need to install Go on your computer.

**\# go run main.go exporter.go postman.go**    
    
Postman file path: /path/to/postman/file.json  
Path to save the file: /path/to/save  
Base struct name: (optional)  
*File saved to: /path/to/save/APIInput.swift*

## API Input Format
```swift
public  struct  APIInput  {
	public  struct  _RequestName_: _BaseStructName_ {
		let  param1: String?
		
		public  var  toJSON: [String: String] {
			var json: [String:  String] = [:]
			if  let param1 = self.param1 {
				json["param1"]  = param1
			}
			return json
		}
		
		public  var  path: String {
			return  "/api_path_url"
		} 
		public var method: String {
			return "POST"
		}
	}
}
```
