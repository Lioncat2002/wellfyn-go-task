import React, { useState } from "react";
import axios from "axios";
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
      <button onClick={() => navigate("/login")}>Login</button>
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

function LoginPage({ handleLogin }) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    handleLogin(email, password);
    navigate("/more-info");
  };

  return (
    <div>
      <h1>Login</h1>
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
        <button type="submit">login</button>
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
    handleUpdateResume(resume);
    navigate("/user-info");
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    setResume(file);
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

  const handleSignup = async (email, password) => {
    await axios
      .post("https://wellfyn-go-task-production.up.railway.app/api/signup", {
        email: email,
        password: password,
      })
      .then((response) => {
        setToken(response.data.token);
        localStorage.setItem("token", response.data.token);
        console.log(response.data);
      })
      .catch((err) => {
        console.log(err);
        axios
          .post("https://wellfyn-go-task-production.up.railway.app/api/login", {
            email: email,
            password: password,
          })
          .then((response) => {
            setToken(response.data.token);
            localStorage.setItem("token", token);
            console.log(response.data);
          })
          .catch((err) => console.log(err));
      });
  };
  const handleLogin = async (email, password) => {
    await axios
      .post("https://wellfyn-go-task-production.up.railway.app/api/login", {
        email: email,
        password: password,
      })
      .then((response) => {
        setToken(response.data.token);
        localStorage.setItem("token", response.data.token);
        console.log(response.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const handleUpdateUser = async (name, phoneNumber, imageUrl) => {
    // Update user logic...
    await setUserInfo({ name, phoneNumber, imageUrl });
    axios
      .post(
        "https://wellfyn-go-task-production.up.railway.app/api/user/update?token=" +
          localStorage.getItem("token"),
        {
          name: name,
          imgurl: imageUrl,
          phno: phoneNumber,
        }
      )
      .then((response) => {
        setUserInfo({ name, phoneNumber, imageUrl });
        console.log(response.data);
      })
      .catch((err) => console.log(err));
  };

  const handleUpdateResume = async (file) => {
    const formData = new FormData();
    formData.append("resume", file);

    // Send the form data using Axios
    const response = await axios.post(
      "https://wellfyn-go-task-production.up.railway.app/api/user/update/resume?token=" +
        localStorage.getItem("token"),
      formData,
      {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      }
    );
    console.log(response);
  };

  return (
    <Router>
      <Routes>
        <Route path="/" element={<SignupPage handleSignup={handleSignup} />} />
        <Route
          path="/login"
          element={<LoginPage handleLogin={handleLogin} />}
        />
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
