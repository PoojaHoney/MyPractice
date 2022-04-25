from fastapi import FastAPI, Depends
#Pydantic is a data validation library that uses some neat tricks 
#from the Python3 type system. 
# FastAPI relies heavily on it to both validate incoming data and 
# serialize outgoing data.
from pydantic import BaseModel
from typing import Optional, List
from sqlalchemy import create_engine
from sqlalchemy.orm import declarative_base, sessionmaker, Session
from sqlalchemy import Boolean, Column, Float, String, Integer

app = FastAPI()

#SQL Alchemy setup
SQLALCHEMY_DATABASE_URL = 'sqlite:///places.db'
# To connect to a database, we need to create a SQLAlchemy engine.
#  The SQLAlchemy engine creates a common interface to the database to 
# execute SQL statements. It does this by wrapping a pool of database 
# connections and a dialect in such a way that they can work together to
#  provide uniform access to the backend database.
engine = create_engine(SQLALCHEMY_DATABASE_URL, echo=True, future=True)
# In order to interact with the database, 
# we need to obtain its handle. A session object is the handle to database
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
#The declarative_base() base class contains a MetaData object where newly 
# defined Table objects are collected. This object is intended to be 
# accessed directly for MetaData -specific operations. 
# Such as, to issue CREATE statements for all tables: 
# engine = create_engine('sqlite://') Base.
Base = declarative_base()

#Dependency
#Creates a method get_db which will be executed whenever you need access to the database
def get_db():
    db = SessionLocal()
    try:
        #yield in Python can be used like the return statement in a function. 
        # When done so, the function instead of returning the output,
        #  it returns a generator that can be iterated upon.
        #  You can then iterate through the generator to extract items
        yield db
    finally:
        db.close()

# A SQLAlchemny ORM Place
class DBPlace(Base):
    __tablename__ = 'places'

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String(50))
    description = Column(String, nullable=True)
    coffee = Column(Boolean)
    wifi = Column(Boolean)
    food = Column(Boolean)
    lat = Column(Float)
    lng = Column(Float)

Base.metadata.create_all(bind=engine)

# A Pydantic Place
class Place(BaseModel):
    name: str
    description: Optional[str] = None
    coffee: bool
    wifi: bool
    food: bool
    lat: float
    lng: float

    class Config:
        #SQLAlchemy is a library that facilitates the communication between Python programs 
        # and databases. 
        # Most of the times, this library is used as an Object Relational Mapper (ORM) 
        # tool that translates Python classes to tables on relational databases 
        # and automatically converts function calls to SQL statements
        orm_mode = True


# Methods for interacting with the database
#db: it’s type is a SQLAlchemy session
def get_place(db: Session, place_id: int):
    return db.query(DBPlace).where(DBPlace.id == place_id).first()


def get_places(db: Session):
    return db.query(DBPlace).all()


def create_place(db: Session, place: Place):
    db_place = DBPlace(**place.dict())
    db.add(db_place)
    db.commit()
    db.refresh(db_place)

    return db_place

# Routes for interacting with the API
@app.post('/places/', response_model=Place)
def create_places_view(place: Place, db: Session = Depends(get_db)):
    db_place = create_place(db, place)
    return db_place


@app.get('/places/', response_model=List[Place])
def get_places_view(db: Session = Depends(get_db)):
    return get_places(db)


@app.get('/place/{place_id}')
def get_place_view(place_id: int, db: Session = Depends(get_db)):
    return get_place(db, place_id) 


#this is a route - It tells FastAPI that the following method should 
# be run when the user requests the / path.
@app.get('/') 
#method declaration
async def root():
    #sends the data to browser
    return {'message': 'Hello Pooja'}