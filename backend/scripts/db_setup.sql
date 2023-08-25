CREATE TABLE IF NOT EXISTS vehicles (
  id char(36) PRIMARY KEY NOT NULL,
  brand varchar(50) NOT NULL,
  model varchar(50) NOT NULL,
  year numeric(4) NOT NULL,
  price numeric(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS leads (
  id char(36) PRIMARY KEY NOT NULL,
  name varchar(50) NOT NULL,
  email varchar(20) NOT NULL,
  phone char(14) NOT NULL,
  chosenVehicleId char(36) NOT NULL
);

INSERT INTO
  vehicles (id, brand, model, year, price)
VALUES 
  ('320d3c04-cc57-435e-92f9-d6e3f6e0a7f4', 'Tesla', 'Model S', 2022, 216000),
  ('1da55e15-f253-488d-8590-e0c8a955898b', 'Ford', 'Mustang', 2021, 372000),
  ('bc169144-f3bd-440f-8f60-5b9c833eda47', 'Ferrari', 'LaFerrari', 2022, 123000);
