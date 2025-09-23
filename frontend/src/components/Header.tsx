import "../App.css"

function Header() {
  return (
    <>
      <div className="bg-gray-500 flex items-center">
        <img src="https://www.svgrepo.com/show/374378/queue.svg" alt="lgo pic" className="w-25 h-25 p-4" />
        <div className=" p-4">
          <p>queue</p>
        </div>
        <div className=" p-4">
          <p>history</p>
        </div>
        <div className=" p-4">
          <p>settings</p>
        </div>
        <div className="flex ml-auto" >
          <div className=" p-4 ">
            <p>dark/light</p>
          </div>
          <div className=" p-4 ">
            <p>login</p>
          </div>
        </div>

      </div>
    </>

  );
}
export default Header;
