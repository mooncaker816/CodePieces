var bytes= 4; 
var aBuffer = new ArrayBuffer(bytes);
var dataView = new DataView(aBuffer);
dataView.setUint32(0, 97);
var textDecoder = new TextDecoder('utf-8');
str = textDecoder.decode(dataView);
console.log(str);