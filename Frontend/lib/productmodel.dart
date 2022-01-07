
class Product{
  String name;
  num price;
  int count;

  Product({required this.name,required this.price,required this.count});
  Product.fromJson(Map<String,dynamic> json)
  : name=json["Name"],
  price=json["Price"],
  count=json["Count"];
}