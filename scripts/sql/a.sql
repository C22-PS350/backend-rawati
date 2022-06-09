DROP DATABASE IF EXISTS rawati;

CREATE DATABASE IF NOT EXISTS rawati;

USE rawati;

CREATE TABLE users (
    user_id INT AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    username VARCHAR(30) NOT NULL,
    email VARCHAR(60) NOT NULL,
    password CHAR(60) NOT NULL,
    is_verified BOOLEAN,
    PRIMARY KEY (user_id),
    UNIQUE (username),
    UNIQUE (email)
);

CREATE TABLE user_token (
    user_id INT,
    token CHAR(40),
    created_at DATETIME NOT NULL,
    PRIMARY KEY (user_id, token),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE (token)
);

CREATE TABLE user_profile (
    profile_id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    gender CHAR(1),
    birth_date DATE,
    height INT,
    weight INT,
    weight_goal INT,
    PRIMARY KEY (profile_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE (user_id)
);

CREATE TABLE exercises (
    exercise_id INT AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    calories DECIMAL(6, 2) NOT NULL,
    PRIMARY KEY (exercise_id)
);

CREATE TABLE foods (
    food_id INT AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    calories DECIMAL(6, 2) NOT NULL,
    PRIMARY KEY (food_id)
);

CREATE TABLE exercise_per_day (
    exercise_activity_id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    name VARCHAR(60) NOT NULL,
    exercise_date DATE NOT NULL,
    duration INT NOT NULL,
    calories DECIMAL(6, 2),
    PRIMARY KEY (exercise_activity_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE food_per_day (
    food_activity_id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    name VARCHAR(60) NOT NULL,
    food_date DATE NOT NULL,
    calories DECIMAL(6, 2) NOT NULL,
    PRIMARY KEY (food_activity_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

INSERT INTO exercises (name, calories) VALUES
    ('Cycling, mountain bike, bmx',646.5),
    ('Cycling, <10 mph, leisure bicycling',304.0),
    ('Cycling, >20 mph, racing',1216.75),
    ('Cycling, 10-11.9 mph, light',456.0),
    ('Cycling, 12-13.9 mph, moderate',608.5),
    ('Cycling, 14-15.9 mph, vigorous',760.5),
    ('Cycling, 16-19 mph, very fast, racing',912.5),
    ('Unicycling',380.25),
    ('Stationary cycling, very light',228.0),
    ('Stationary cycling, light',418.25),
    ('Stationary cycling, moderate',532.25),
    ('Stationary cycling, vigorous',798.5),
    ('Stationary cycling, very vigorous',950.75),
    ('Calisthenics, vigorous, pushups, situps…',608.5),
    ('Calisthenics, light',266.25),
    ('Circuit training, minimal rest',608.5),
    ('Weight lifting, body building, vigorous',456.0),
    ('Weight lifting, light workout',228.0),
    ('Health club exercise',418.25),
    ('Stair machine',684.25),
    ('Rowing machine, light',266.25),
    ('Rowing machine, moderate',532.25),
    ('Rowing machine, vigorous',646.5),
    ('Rowing machine, very vigorous',912.5),
    ('Ski machine',532.25),
    ('Aerobics, low impact',380.25),
    ('Aerobics, high impact',532.25),
    ('Aerobics, step aerobics',646.5),
    ('Aerobics, general',494.25),
    ('Jazzercise',456.0),
    ('Stretching, hatha yoga',304.0),
    ('Mild stretching',190.25),
    ('Instructing aerobic class',456.0),
    ('Water aerobics',304.0),
    ('Ballet, twist, jazz, tap',342.5),
    ('Ballroom dancing, slow',228.0),
    ('Ballroom dancing, fast',418.25),
    ('Running, 5 mph (12 minute mile)',608.5),
    ('Running, 5.2 mph (11.5 minute mile)',684.25),
    ('Running, 6 mph (10 min mile)',760.5),
    ('Running, 6.7 mph (9 min mile)',836.5),
    ('Running, 7 mph (8.5 min mile)',874.5),
    ('Running, 7.5mph (8 min mile)',950.75),
    ('Running, 8 mph (7.5 min mile)',1026.5),
    ('Running, 8.6 mph (7 min mile)',1064.5),
    ('Running, 9 mph (6.5 min mile)',1140.75),
    ('Running, 10 mph (6 min mile)',1216.75),
    ('Running, 10.9 mph (5.5 min mile)',1368.75),
    ('Running, cross country',684.25),
    ('Running, general',608.5),
    ('Running, on a track, team practice',760.5),
    ('Running, stairs, up',1140.75),
    ('Track and field (shot, discus)',304.0),
    ('Track and field (high jump, pole vault)',456.0),
    ('Track and field (hurdles)',760.5),
    ('Archery',266.25),
    ('Badminton',342.5),
    ('Basketball game, competitive',608.5),
    ('Playing basketball, non game',456.0),
    ('Basketball, officiating',532.25),
    ('Basketball, shooting baskets',342.5),
    ('Basketball, wheelchair',494.25),
    ('Running, training, pushing wheelchair',608.5),
    ('Billiards',190.25),
    ('Bowling',228.0),
    ('Boxing, in ring',912.5),
    ('Boxing, punching bag',456.0),
    ('Boxing, sparring',684.25),
    ('Coaching: football, basketball, soccer…',304.0),
    ('Cricket (batting, bowling)',380.25),
    ('Croquet',190.25),
    ('Curling',304.0),
    ('Darts (wall or lawn)',190.25),
    ('Fencing',456.0),
    ('Football, competitive',684.25),
    ('Football, touch, flag, general',608.5),
    ('Football or baseball, playing catch',190.25),
    ('Frisbee playing, general',228.0),
    ('Frisbee, ultimate frisbee',608.5),
    ('Golf, general',342.5),
    ('Golf, walking and carrying clubs',342.5),
    ('Golf, driving range',228.0),
    ('Golf, miniature golf',228.0),
    ('Golf, walking and pulling clubs',327.0),
    ('Golf, using power cart',266.25),
    ('Gymnastics',304.0),
    ('Hacky sack',304.0),
    ('Handball',912.5),
    ('Handball, team',608.5),
    ('Hockey, field hockey',608.5),
    ('Hockey, ice hockey',608.5),
    ('Riding a horse, general',304.0),
    ('Horesback riding, saddling horse',266.25),
    ('Horseback riding, grooming horse',266.25),
    ('Horseback riding, trotting',494.25),
    ('Horseback riding, walking',190.25),
    ('Horse racing, galloping',608.5),
    ('Horse grooming, moderate',456.0),
    ('Horseshoe pitching',228.0),
    ('Jai alai',912.5),
    ('Martial arts, judo, karate, jujitsu',760.5),
    ('Martial arts, kick boxing',760.5),
    ('Martial arts, tae kwan do',760.5),
    ('Krav maga training',760.5),
    ('Juggling',304.0),
    ('Kickball',532.25),
    ('Lacrosse',608.5),
    ('Orienteering',684.25),
    ('Playing paddleball',456.0),
    ('Paddleball, competitive',760.5),
    ('Polo',608.5),
    ('Racquetball, competitive',760.5),
    ('Playing racquetball',532.25),
    ('Rock climbing, ascending rock',836.5),
    ('Rock climbing, rappelling',608.5),
    ('Jumping rope, fast',912.5),
    ('Jumping rope, moderate',760.5),
    ('Jumping rope, slow',608.5),
    ('Rugby',760.5),
    ('Shuffleboard, lawn bowling',228.0),
    ('Skateboarding',380.25),
    ('Roller skating',532.25),
    ('Roller blading, in-line skating',912.5),
    ('Sky diving',228.0),
    ('Soccer, competitive',760.5),
    ('Playing soccer',532.25),
    ('Softball or baseball',380.25),
    ('Softball, officiating',304.0),
    ('Softball, pitching',456.0),
    ('Squash',912.5),
    ('Table tennis, ping pong',304.0),
    ('Tai chi',304.0),
    ('Playing tennis',532.25),
    ('Tennis, doubles',456.0),
    ('Tennis, singles',608.5),
    ('Trampoline',266.25),
    ('Volleyball, competitive',608.5),
    ('Playing volleyball',228.0),
    ('Volleyball, beach',608.5),
    ('Wrestling',456.0);

INSERT INTO foods (name, calories) VALUES
    ('Cornstarch',381),
    ('Nuts, pecans',691),
    ('Teff, uncooked',367),
    ('Sherbet, orange',144),
    ('Cheese, camembert',300),
    ('Vegetarian fillets',290),
    ('Goji berries, dried',349),
    ('Mango nectar, canned',51),
    ('Crackers, rusk toast',407),
    ('Chicken, boiled, feet',215),
    ('Pie, lemon, fried pies',316),
    ('Salami, turkey, cooked',172),
    ('Spices, ground, savory',272),
    ('Candies, sesame crunch',516),
    ('Cheese, low fat, cream',201),
    ('Syrup, Canadian, maple',270),
    ('Chewing gum, sugarless',268),
    ('Nuts, dried, pine nuts',673),
    ('Pasta, unenriched, dry',371),
    ('Cookies, Marie biscuit',406),
    ('Nuts, dried, beechnuts',576),
    ('Currants, dried, zante',283),
    ('Gravy, mix, dry, onion',322),
    ('Pie, fruit, fried pies',316),
    ('Snacks, cakes, popcorn',384),
    ('Snack, Mixed Berry Bar',383),
    ('Babyfood, pear, juice',47),
    ('Broccoli raab, cooked',33),
    ('Butter oil, anhydrous',876),
    ('Egg custards, dry mix',410),
    ('Peanut flour, low fat',428),
    ('Fish, smoked, haddock',116),
    ('Ground turkey, cooked',203),
    ('Bread, toasted, wheat',313),
    ('Danish pastry, cheese',374),
    ('Nuts, glazed, walnuts',500),
    ('Spices, garlic powder',331),
    ('Oil, soybean lecithin',763),
    ('Beef, pastrami, cured',147),
    ('Frankfurter, meatless',233),
    ('Ice creams, chocolate',216),
    ('Snacks, potato sticks',522),
    ('Figs, uncooked, dried',249),
    ('Syrup, fruit flavored',261),
    ('Cream, cultured, sour',198),
    ('Gravy, canned, au jus',16),
    ('Cheese, port de salut',352),
    ('Soup, mix, dry, onion',293),
    ('Bacon and beef sticks',517),
    ('Salami, pork, Italian',425),
    ('Crackers, whole-wheat',427),
    ('Hominy, white, canned',72),
    ('Horseradish, prepared',48),
    ('Lebanon bologna, beef',172),
    ('Ham and cheese spread',245),
    ('Sauce, worcestershire',78),
    ('Candies, marshmallows',318),
    ('Hummus, home prepared',177),
    ('Horned melon (Kiwano)',44),
    ('Ice cream sundae cone',254),
    ('Cabbage, cooked, napa',12),
    ('Peppers, dried, ancho',281),
    ('Parsley, freeze-dried',271),
    ('Spices, black, pepper',251),
    ('Spices, white, pepper',296),
    ('Nuts, dried, pilinuts',719),
    ('Candies, butterscotch',391),
    ('Potato salad with egg',157),
    ('Cheese, Mexican blend',358),
    ('Papaya nectar, canned',57),
    ('Yeast extract spread',185),
    ('Pasta, enriched, dry',371),
    ('Cookies, gingersnaps',416),
    ('Sour cream, fat free',74),
    ('Frankfurter, chicken',223),
    ('Ham, canned, chopped',239),
    ('Spices, dried, thyme',276),
    ('Carrot juice, canned',40),
    ('Corn, dried (Navajo)',386),
    ('Pate, truffle flavor',327),
    ('Fruit butters, apple',173),
    ('Fast foods, coleslaw',153),
    ('Salami, beef, cooked',261),
    ('Spices, chili powder',282),
    ('Spices, dried, basil',233),
    ('Pickle relish, sweet',130),
    ('Chives, freeze-dried',311),
    ('Crackers, multigrain',482),
    ('Spices, ground, mace',475),
    ('Spices, onion powder',341),
    ('Butter, without salt',717),
    ('Rice noodles, cooked',108),
    ('Barley flour or meal',345),
    ('Garlic bread, frozen',350),
    ('Rolls, sweet, dinner',321),
    ('Snacks, banana chips',519),
    ('Rolls, wheat, dinner',273),
    ('Candies, HEATH BITES',530),
    ('Seaweed, dried, agar',306),
    ('Bread, cracked-wheat',260),
    ('Polish sausage, pork',326),
    ('Spices, curry powder',325),
    ('Bacon bits, meatless',476),
    ('Semolina, unenriched',360),
    ('Grape leaves, canned',69),
    ('Spices, ground, sage',315),
    ('Oil, corn and canola',884),
    ('Celery flakes, dried',319),
    ('Crackers, egg, matzo',391),
    ('Loganberries, frozen',55),
    ('Soybean, curd cheese',151),
    ('Tostada shells, corn',474),
    ('Fish, smoked, cisco',177),
    ('Frankfurter, turkey',223),
    ('Candies, jellybeans',375),
    ('Oil, ucuhuba butter',884),
    ('Soy protein isolate',335),
    ('Syrups, light, corn',283),
    ('Mother''s loaf, pork',282),
    ('Tapioca, dry, pearl',358),
    ('Spices, celery seed',392),
    ('Candies, peanut bar',522),
    ('Bread, toasted, rye',284),
    ('Oil, apricot kernel',884),
    ('Candies, YORK BITES',394),
    ('Meatballs, meatless',197),
    ('Sweet rolls, cheese',360),
    ('Figs, stewed, dried',107),
    ('Bread sticks, plain',412),
    ('Milk and cereal bar',413),
    ('Bread, pumpernickel',250),
    ('Prune juice, canned',71),
    ('Soy flour, defatted',327),
    ('Bread, toasted, egg',315),
    ('Spices, fennel seed',345),
    ('Toppings, pineapple',253),
    ('Acorn stew (Apache)',95),
    ('Tomatoes, sun-dried',258),
    ('Guava sauce, cooked',36),
    ('Melon balls, frozen',33),
    ('Fish oil, cod liver',902),
    ('Nuts, dried, acorns',509),
    ('Rolls, pumpernickel',276),
    ('Gravy, dry, chicken',381),
    ('Ice creams, vanilla',207),
    ('Corn grain, yellow',365),
    ('Ice cream sandwich',237),
    ('Croutons, seasoned',465),
    ('Snacks, taro chips',498),
    ('Liver cheese, pork',304),
    ('Jams and preserves',278),
    ('Oil, nutmeg butter',884),
    ('Semolina, enriched',360),
    ('Spices, anise seed',337),
    ('Snacks, corn cakes',387),
    ('Catsup, low sodium',101),
    ('Croissants, cheese',414),
    ('Milk, fluid, sheep',108),
    ('Fast food, biscuit',370),
    ('Bagels, multigrain',241),
    ('Beef, dried, cured',153),
    ('Nuts, almond paste',458),
    ('Fish oil, menhaden',902),
    ('Rolls, egg, dinner',307),
    ('Bread, white wheat',238),
    ('Rolls, rye, dinner',286),
    ('Gravy, dry, au jus',313),
    ('Cheese, neufchatel',253),
    ('Hummus, commercial',166),
    ('Croissants, butter',406),
    ('Carrot, dehydrated',341),
    ('Vinegar, distilled',18),
    ('Soy flour, low-fat',372),
    ('Syrups, sugar free',52),
    ('Syrups, dark, corn',286),
    ('Taco shells, baked',476),
    ('Spices, cumin seed',375),
    ('Whey, dried, sweet',353),
    ('Whey, fluid, sweet',27),
    ('Spices, poppy seed',525),
    ('Babyfood, pretzels',397),
    ('Fat, mutton tallow',902),
    ('Coffeecake, cheese',339),
    ('Sauce, horseradish',503),
    ('Dates, deglet noor',282),
    ('Gravy, dry, turkey',367),
    ('Vital wheat gluten',370),
    ('Sugars, granulated',387),
    ('Frankfurter, pork',269),
    ('Cookies, molasses',430),
    ('Cookies, fig bars',348),
    ('Liverwurst spread',305),
    ('Roast beef spread',223),
    ('Sausage, meatless',255),
    ('Frankfurter, meat',290),
    ('Rice flour, brown',363),
    ('Peppermint, fresh',70),
    ('Vinegar, red wine',19),
    ('Vinegar, balsamic',88),
    ('Luxury loaf, pork',141),
    ('Rye flour, medium',349),
    ('Bread, wheat bran',248),
    ('Croissants, apple',254),
    ('Candies, Tamarind',368),
    ('Barley malt flour',361),
    ('Rice noodles, dry',364),
    ('Fish oil, sardine',902),
    ('Fish oil, herring',902),
    ('Plantains, cooked',116),
    ('Butterbur, canned',3),
    ('Gravy, dry, brown',367),
    ('Cheese, provolone',351),
    ('Cheese, roquefort',369),
    ('Muffins, oat bran',270),
    ('Syrups, grenadine',268),
    ('Raisins, seedless',299),
    ('Wheat germ, crude',360),
    ('Wheat, hard white',342),
    ('Wheat bran, crude',216),
    ('Spices, dill seed',305),
    ('Oil, cocoa butter',884),
    ('Marmalade, orange',246),
    ('Whey, dried, acid',339),
    ('Whey, fluid, acid',24),
    ('Eggplant, pickled',49),
    ('Babyfood, cookies',433),
    ('Egg, dried, whole',592),
    ('Coffeecake, fruit',311),
    ('Candies, caramels',382),
    ('Sour cream, light',136),
    ('Egg, dried, white',382),
    ('Wild rice, cooked',101),
    ('Chicken, meatless',224),
    ('Corn grain, white',365),
    ('Wheat, soft white',340),
    ('Cheese, limburger',327),
    ('Corn bran, crude',224),
    ('Bread, rice bran',243),
    ('Sugars, powdered',389),
    ('Cookies, fortune',378),
    ('Couscous, cooked',112),
    ('Dill weed, fresh',43),
    ('Spearmint, dried',285),
    ('Spearmint, fresh',44),
    ('Tamales (Navajo)',153),
    ('Rice bran, crude',316),
    ('Rye flour, light',357),
    ('Spices, bay leaf',313),
    ('Spices, cardamom',311),
    ('Fish oil, salmon',902),
    ('Fish, tuna salad',187),
    ('Cheese, muenster',368),
    ('Cheese, cheshire',387),
    ('Bagels, oat bran',255),
    ('Fat, beef tallow',902),
    ('Oat bran, cooked',40),
    ('Sugar, turbinado',399),
    ('Headcheese, pork',157),
    ('Ham salad spread',216),
    ('Olive loaf, pork',235),
    ('Pastrami, turkey',139),
    ('Egg, dried, yolk',669),
    ('Quinoa, uncooked',368),
    ('Cheese, monterey',373),
    ('Cheese, cheddar',404),
    ('Oil, wheat germ',884),
    ('Bread, oat bran',236),
    ('Potato pancakes',268),
    ('Spinach souffle',172),
    ('Bacon, meatless',309),
    ('Rosemary, fresh',131),
    ('Syrups, sorghum',290),
    ('Rye flour, dark',325),
    ('Oil, tomatoseed',884),
    ('Croutons, plain',407),
    ('Spices, paprika',282),
    ('Sauce, barbecue',172),
    ('Arrowroot flour',357),
    ('Pimento, canned',23),
    ('Raisins, seeded',296),
    ('Bologna, turkey',209),
    ('Vanilla extract',288),
    ('Wheat, sprouted',198),
    ('Bread, cinnamon',253),
    ('Spices, saffron',310),
    ('Spelt, uncooked',338),
    ('Cabbage, kimchi',15),
    ('Seeds, flaxseed',534),
    ('Cheese, gruyere',413),
    ('Cheese, gjetost',466),
    ('Cheese, fontina',389),
    ('Butter, salted',717),
    ('Oil, rice bran',884),
    ('Crackers, milk',446),
    ('Dates, medjool',277),
    ('Bread, oatmeal',269),
    ('Bread, italian',271),
    ('Millet, puffed',354),
    ('Vinegar, cider',21),
    ('Scrapple, pork',213),
    ('Oil, poppyseed',884),
    ('Oil, grapeseed',884),
    ('Longans, dried',286),
    ('Pectin, liquid',11),
    ('Pretzels, soft',338),
    ('Quinoa, cooked',120),
    ('Parsley, fresh',36),
    ('Barley, hulled',354),
    ('Cheese, romano',387),
    ('Dulce de Leche',315),
    ('Bulgur, cooked',83),
    ('Tempeh, cooked',195),
    ('Chicken spread',158),
    ('Millet, cooked',119),
    ('Capers, canned',23),
    ('Litchis, dried',277),
    ('Strudel, apple',274),
    ('Oil, cupu assu',884),
    ('Cheese, tilsit',340),
    ('Cheese, brick',371),
    ('Sugars, brown',380),
    ('Candied fruit',322),
    ('Bologna, beef',299),
    ('Blood sausage',379),
    ('Nuts, almonds',579),
    ('Couscous, dry',376),
    ('Candies, hard',394),
    ('Meat extender',311),
    ('Syrups, maple',260),
    ('Sorghum grain',329),
    ('Oil, hazelnut',884),
    ('Pepeao, dried',298),
    ('Rolls, french',277),
    ('Bagels, wheat',250),
    ('Sugars, maple',354),
    ('Rice crackers',416),
    ('Cheese, colby',394),
    ('Cheese, cream',350),
    ('Bologna, pork',247),
    ('Spelt, cooked',127),
    ('Bread, cheese',408),
    ('Bread, potato',266),
    ('Cheese, swiss',393),
    ('Cheese, gouda',356),
    ('Tomato powder',302),
    ('Cheese, blue',353),
    ('Cheese, brie',334),
    ('Thyme, fresh',101),
    ('Basil, fresh',23),
    ('Syrups, malt',318),
    ('Oil, teaseed',884),
    ('Oil, sheanut',884),
    ('Oil, babassu',884),
    ('Fish, surimi',99),
    ('Bread, wheat',267),
    ('Cheese, edam',357),
    ('Cheese, feta',264),
    ('Cracker meal',383),
    ('Oil, coconut',892),
    ('Wheat, durum',339),
    ('Millet flour',382),
    ('Potato flour',357),
    ('Teff, cooked',101),
    ('Oil, mustard',884),
    ('Oil, avocado',884),
    ('Phyllo dough',299),
    ('Fat, chicken',900),
    ('Fruit syrup',341),
    ('Carob flour',222),
    ('Prune puree',257),
    ('Salt, table',0),
    ('Chewing gum',360),
    ('Fat, turkey',900),
    ('Tofu, fried',270),
    ('Tofu yogurt',94),
    ('Bagels, egg',278),
    ('Bulgur, dry',342),
    ('Ham, minced',263),
    ('Oil, almond',884),
    ('Oil, walnut',884),
    ('Oil, canola',884),
    ('Syrup, Cane',269),
    ('Clif Z bar',409),
    ('Fish broth',16),
    ('Fat, goose',900),
    ('Bread, rye',259),
    ('Pie, peach',224),
    ('Bread, egg',287),
    ('Buckwheat',343),
    ('Oil, palm',884),
    ('Fat, duck',882),
    ('Triticale',336),
    ('Rye grain',338),
    ('Molasses',290),
    ('Zwieback',426),
    ('Oil, oat',884),
    ('Jellies',266),
    ('Eggnog',88),
    ('Catsup',101),
    ('Tempeh',192),
    ('Wasabi',292),
    ('Honey',304),
    ('Natto',211),
    ('Okara',76),
    ('Papad',371),
    ('Lard',902),
    ('Oats',389),
    ('Miso',198),
    ('Poi',112);
