# PostmanToAPIInput
My first code using Go. Converting Postman JSON file to APIInput.swift. As this is my first Go codes, so do not expect a good codes.

## How to use
Make sure you have installed Go on your computer.  
### **\# go run main.go**
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
