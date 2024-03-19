CREATE TABLE product (
  id VARCHAR(255),
  name VARCHAR(255),
  PRIMARY KEY(id)
);

CREATE TABLE composition (
  product_id VARCHAR(255),
  material_name VARCHAR(255),
  material_price_per_unit VARCHAR(255),
  unit VARCHAR(255),
  material_quantity VARCHAR(255),
  currency VARCHAR(255),
  manufacture_date VARCHAR(255),
  CONSTRAINT fk_product_id
  	FOREIGN KEY(product_id)
		REFERENCES product(id)
);
