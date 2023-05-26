import React, { useState } from "react";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  Navigate,
  useNavigate,
} from "react-router-dom";

function SignupPage({ handleSignup }) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    handleSignup(email, password);
    navigate("/more-info");
  };

  return (
    <div>
      <h1>Signup</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Signup</button>
      </form>
    </div>
  );
}

function MoreInfoPage({ handleUpdateUser, handleUpdateResume }) {
  const [name, setName] = useState("");
  const [phoneNumber, setPhoneNumber] = useState("");
  const [imageUrl, setImageUrl] = useState("");
  const [resume, setResume] = useState(null);
  const navigate = useNavigate();
  const handleSubmit = (e) => {
    e.preventDefault();
    handleUpdateUser(name, phoneNumber, imageUrl);
    navigate("/user-info");
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    setResume(file);
    handleUpdateResume(file);
  };

  return (
    <div>
      <h1>More Info</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="tel"
          placeholder="Phone Number"
          value={phoneNumber}
          onChange={(e) => setPhoneNumber(e.target.value)}
        />
        <input
          type="text"
          placeholder="Image URL"
          value={imageUrl}
          onChange={(e) => setImageUrl(e.target.value)}
        />
        <input type="file" onChange={handleFileChange} />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
}

function UserInfoDisplay({ name, phoneNumber, imageUrl }) {
  return (
    <div>
      <h1>User Info</h1>
      <p>Name: {name}</p>
      <p>Phone Number: {phoneNumber}</p>
      <img src={imageUrl} alt="User" />
    </div>
  );
}

function App() {
  const [token, setToken] = useState("");
  const [userInfo, setUserInfo] = useState({
    name: "",
    phoneNumber: "",
    imageUrl: "",
  });

  const handleSignup = (email, password) => {
    // Signup logic...

    // Assuming successful signup, set token and store in localStorage
    const token = "your-jwt-token";
    setToken(token);
    localStorage.setItem("token", token);
  };

  const handleUpdateUser = (name, phoneNumber, imageUrl) => {
    // Update user logic...
    setUserInfo({ name, phoneNumber, imageUrl });
    console.log(name, phoneNumber, imageUrl);
    console.log(token);
  };

  const handleUpdateResume = (file) => {
    // Update resume logic...
  };

  return (
    <Router>
      <Routes>
        <Route path="/" element={<SignupPage handleSignup={handleSignup} />} />
        <Route
          path="/more-info"
          element={
            localStorage.getItem("token") ? (
              <MoreInfoPage
                handleUpdateUser={handleUpdateUser}
                handleUpdateResume={handleUpdateResume}
              />
            ) : (
              <Navigate to="/" replace /> // Redirect to "/" if not authenticated
            )
          }
        />
        <Route path="/user-info" element={<UserInfoDisplay {...userInfo} />} />
      </Routes>
    </Router>
  );
}

export default App;
