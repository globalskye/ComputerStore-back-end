import React, { useState, useEffect, Fragment } from "react";
import { Col, Row } from "react-bootstrap";
import { Button, Card } from "react-bootstrap";
import {
  MDBRadio,
  MDBContainer,
  MDBRow,
  MDBCol,
  MDBCard,
  MDBCardBody,
  MDBCardImage,
  MDBIcon,
  MDBRipple,
  MDBBtn,
  MDBBtnGroup,
} from "mdb-react-ui-kit";

import { getAllCategories, getAllProductItems } from "../services/store.service";
type Item = {
  id: number;
  name: string;
  image: string;
  description: string;
  price: number;
  garantia:string;
  category:string;
};
type Category = {
  category:string;
}
const Home: React.FC = () => {
  const [products, setProduct] = useState<Item[]>();
  const [categories, setCategory] = useState<Category[]>();
  const [isToggleOn, setToggle] = useState<boolean>();

  

  useEffect(() => {
    getAllProductItems().then(     
      (response) => {
        setProduct(response.data);
      },
      (error) => {
        const _content =
          (error.response && error.response.data) ||
          error.message ||
          error.toString();

        setProduct(_content);
      }
    );
    getAllCategories().then(
      (response) => {
        setCategory(response.data);
      },
      (error) => {
        const _content =
          (error.response && error.response.data) ||
          error.message ||
          error.toString();

        setCategory(_content);
      }
    )
  }, []);
 

  function categoriesList() {
    if (categories) {
      return (
        <>
        <MDBCol md="3" xl="3">
          <MDBCard className=" border rounded-3 mt-5 mb-3">
                <h5>
                  Categories
                </h5>
                {/* <div className="form-check">
  <input className="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault" checked/>
  <div>
    All
  </div>
</div> */}
<div className="form-check">
              <input className="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault1"checked/>
              <div>
              All
              </div>
              </div>
                {categories.map((item:Category) => 
               (
              <div className="form-check">
              <input className="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault1"/>
              <div>
              {item.category}
              </div>
              </div>
              
              ))}
              
                
              </MDBCard>
            </MDBCol>
        </>
      )
    }
  }
  function itemLists() {
    if (products) {
      return (
      <>
      {products.map((item:Item) =>
      (
        <MDBCardBody>
            <MDBRow>
              <MDBCol md="12" lg="3" className="mb-4 mb-lg-0">
                <MDBRipple
                  rippleColor="light"
                  rippleTag="div"
                  className="bg-image rounded hover-zoom hover-overlay"
                >
                  <MDBCardImage
                    src={item.image}
                    fluid
                    className="w-100"
                  />
                  <a href="#!">
                    <div
                      className="mask"
                      style={{ backgroundColor: "rgba(251, 251, 251, 0.15)" }}
                    ></div>
                  </a>
                </MDBRipple>
              </MDBCol>
              <MDBCol md="6">
                <h5>{item.name}</h5>
                
                  {item.description}
              </MDBCol>
              <MDBCol
                md="6"
                lg="3"
                className="border-sm-start-none border-start"
              >
                <div className="d-flex flex-row align-items-center mb-1">
                  <h4 className="mb-1 me-1">${item.price}</h4>
                </div>
                <h6 className="text-success">Гарантия {item.garantia} месяцев</h6>
                <div className="d-flex flex-column mt-4">
                  <MDBBtn color="primary" size="sm">
                    Details
                  </MDBBtn>
                  <MDBBtn outline color="primary" size="sm" className="mt-2">
                    Add to card
                  </MDBBtn>
                </div>
              </MDBCol>
            </MDBRow>
          </MDBCardBody>
      ))}
      </>)
    }
  }
  


  

  return (
    <>
    
    
    <MDBContainer >
      
    <MDBRow className="justify-content-center; mb-0">
    {categoriesList()}
      <MDBCol md="8" xl="9">
        <MDBCard className=" border rounded-3 mt-5 mb-3">
          
          {itemLists()}
        </MDBCard>
      </MDBCol>
    </MDBRow>
  </MDBContainer>
  </>
  );
//   return (
//     <body className="background">
//   <div className="center">
   
//   </div>
// </body>
//   );
};

export default Home;
