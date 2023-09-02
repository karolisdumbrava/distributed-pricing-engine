import { useEffect, useState } from "react";
import { getRandomLocation } from "./api/getRandomLocation";
import { getProductsPricesByLocation } from "./api/getProductsPricesByLocation";

function App() {
  const [location, setLocation] = useState(null);
  const [products, setProducts] = useState([]);
  console.log(location)

  useEffect(() => {
    const fetchLocation = async () => {
      try {
        const location = await getRandomLocation();
        setLocation(location);
      } catch (error) {
        console.log(error);
    }
  }
    fetchLocation();
  }, []);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const products = await getProductsPricesByLocation(location.id);
        setProducts(products);
      } catch (error) {
        console.log(error);
      }
    }
    if(location && location.id){
      fetchProducts();
    }
  }, [location])

  return (
    <div>
      <div>
        {location ? (
          <div>
            Current location: 
            <br />
            {location.city}, {location.country}
          </div>
        ) : (
          <div>Loading...</div>
        )}
      </div>
      <div>
        {products.length > 0 ? (
          <div>
            <h2>Products</h2>
            <ul>
              {products.map((product) => (
                <li key={product.id}>
                  {product.name} - {product.final_price} {product.currency}
                </li>
              ))}
            </ul>
          </div>
        ) : (
          <div>Loading...</div>
        )}
      </div>
    </div>
  );
}

export default App;