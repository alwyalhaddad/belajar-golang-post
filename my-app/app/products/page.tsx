import CardList from "../components/Products/CardList";
import ViewUserButton from "../components/Products/ViewUserButton";

interface Iproduct {
  id: number;
  name: string;
  description: string;
  price: number;
  stock_quantity: number;
}

const Products = async () => {
  const apiUrl =
    process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/products";
  const response = await fetch(apiUrl);

  if (!response.ok) {
    console.error(`HTTP error! status: ${response.status}`);
    const text = await response.text();
    console.log("Server resposne:", text);
    return;
  }

  // Get object response from API
  const apiResponse = await response.json();

  // Get array product from property 'products'
  const products: Iproduct[] = apiResponse.products;
  return (
    <>
      <h1 className="text-fuchsia-500">Products Page</h1>
      {products &&
        products.map((product) => {
          <CardList key={product.id}>
            <p>ID: {product.id}</p>
            <p>Name: {product.name}</p>
            <p>Description: {product.description}</p>
          </CardList>;
        })}
    </>
  );
};

export default Products;
