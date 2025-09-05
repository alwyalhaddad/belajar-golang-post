"use client";

const ViewUserButton = () => {
  const handleClick = () => alert("CLICK CLICK CLICK");

  return (
    <>
      <button onClick={handleClick}>Lihat User</button>
    </>
  );
};

export default ViewUserButton;
