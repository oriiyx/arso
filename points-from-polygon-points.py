from shapely.geometry import Point, Polygon
import random

# Define the polygon using the given latitude and longitude coordinates
polygon_points = [
     {
         45.31795190201766,
         16.66718753586094
       },
       {
         45.36262446643174,
         16.68514563524353
       },
       {
         45.38952242370733,
         16.68048635175977
       },
       {
         45.40197523066536,
         16.66856909767565
       },
       {
         45.42138348422668,
         16.64588986728298
       },
       {
         45.43316211181928,
         16.61831458123825
       },
       {
         45.44426043022825,
         16.5779103713106
       },
       {
         45.42704000790972,
         16.51587484404468
       },
       {
         45.38923815401012,
         16.47462090992648
       },
       {
         45.34096150407625,
         16.484019527811
       },
       {
         45.30572941731185,
         16.51348982762373
       },
       {
         45.29690782846773,
         16.54846732839304
       },
       {
         45.29021330706069,
         16.61501539405956
       },
       {
         45.31795190201766,
         16.66718753586094
       }
    ]

# Create a polygon
polygon = Polygon(polygon_points)

# Generate random points within the bounding box of the polygon
points_inside = []
while len(points_inside) < 5:
    minx, miny, maxx, maxy = polygon.bounds
    random_point = Point(random.uniform(minx, maxx), random.uniform(miny, maxy))
    if polygon.contains(random_point):
        points_inside.append(random_point)

# Print the points
for point in points_inside:
    print(f"Longitude: {point.x}, Latitude: {point.y}")

