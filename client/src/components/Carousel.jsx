import Slider from "react-slick";

const Carousel = () => {
  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
  };
  return (
    <div style={{ width: "100%", maxWidth: "1450px", margin: "0 auto" }}>
      <Slider {...settings}>
        <div></div>
        <div>
          <h3>2</h3>
        </div>
      </Slider>
    </div>
  );
};

export default Carousel;
