INSERT INTO m_restaurants ("name",address,created_at,created_by,updated_at,updated_by,deleted_at,deleted_by) VALUES
	 ('Hanamasa','Jalan Mangga 1','2022-11-04 15:47:47.514304+07',NULL,NULL,NULL,NULL,NULL),
	 ('Yoshinoya','Jalan Manggis 29','2022-11-04 15:47:47.522173+07',NULL,NULL,NULL,NULL,NULL),
	 ('KFC','Jalan Kedondong 89','2022-11-04 15:48:07.238516+07',NULL,NULL,NULL,NULL,NULL);

INSERT INTO m_dishes (restaurant_id,"name",price,created_at,created_by,updated_at,updated_by,deleted_at,deleted_by) VALUES
	 (1,'Tenderloin Wagyu 100gr',1890000.0,'2022-11-04 15:48:30.686845+07',NULL,NULL,NULL,NULL,NULL),
	 (2,'Beef Bowl S',45000.0,'2022-11-04 15:49:31.574032+07',NULL,NULL,NULL,NULL,NULL),
	 (2,'Chicken Teriyaki',35000.0,'2022-11-04 15:49:31.581291+07',NULL,NULL,NULL,NULL,NULL),
	 (3,'Spagheti',20000.0,'2022-11-04 15:49:31.588392+07',NULL,NULL,NULL,NULL,NULL),
	 (3,'Large Chicken Crispy',16000.0,'2022-11-04 15:50:02.806566+07',NULL,NULL,NULL,NULL,NULL),
	 (3,'Large Chicken Original',16500.0,'2022-11-04 15:50:02.812349+07',NULL,NULL,NULL,NULL,NULL),
	 (1,'All You Can Eat',200000.0,'2022-11-04 15:50:20.81462+07',NULL,NULL,NULL,NULL,NULL);
