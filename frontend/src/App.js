import { useEffect, useState } from "react";
import { getRandomLocation } from "./api/getRandomLocation";
import { getProductsPricesByLocation } from "./api/getProductsPricesByLocation";

function App() {
  const [location, setLocation] = useState(null);
  const [products, setProducts] = useState([]);
  console.log(location)
  console.log("App rendered")

  useEffect(() => {
    const fetchLocationAndProducts = async () => {
      try {
        const location = await getRandomLocation();
        setLocation(location);
        const products = await getProductsPricesByLocation(location.id);
        setProducts(products);
      } catch (error) {
        console.log(error);
      }
    }
    fetchLocationAndProducts();
  }, []);

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
                  {product.name} - {product.final_price} {location.currency}
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