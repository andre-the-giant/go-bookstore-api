const API_BASE = '/api';

function App() {
  const [token, setToken] = React.useState(sessionStorage.getItem('token'));
  const [books, setBooks] = React.useState([]);
  const [tab, setTab] = React.useState('login');
  const [loginData, setLoginData] = React.useState({ email: '', password: '' });
  const [regData, setRegData] = React.useState({ email: '', password: '' });

  React.useEffect(() => {
    if (token) {
      fetch(`${API_BASE}/books`, {
        headers: { Authorization: `Bearer ${token}` },
      })
        .then((res) => res.json())
        .then(setBooks)
        .catch(console.error);
    }
  }, [token]);

  function handleLogin(e) {
    e.preventDefault();
    fetch(`${API_BASE}/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(loginData),
    })
      .then((res) => res.json())
      .then((data) => {
        if (data.token) {
          sessionStorage.setItem('token', data.token);
          setToken(data.token);
        }
      })
      .catch(console.error);
  }

  function handleRegister(e) {
    e.preventDefault();
    fetch(`${API_BASE}/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(regData),
    })
      .then((res) => res.json())
      .then(() => {
        setTab('login');
      })
      .catch(console.error);
  }

  function logout() {
    sessionStorage.removeItem('token');
    setToken(null);
    setBooks([]);
  }

  if (!token) {
    return (
      <div className="max-w-md mx-auto mt-10 bg-white p-4 shadow">
        <div className="flex mb-4 border-b">
          <button
            className={`flex-1 p-2 ${tab === 'login' ? 'border-b-2 border-blue-500' : ''}`}
            onClick={() => setTab('login')}
          >
            Log In
          </button>
          <button
            className={`flex-1 p-2 ${tab === 'register' ? 'border-b-2 border-blue-500' : ''}`}
            onClick={() => setTab('register')}
          >
            Register
          </button>
        </div>
        {tab === 'login' && (
          <form onSubmit={handleLogin} className="flex flex-col gap-2">
            <input
              type="email"
              className="border p-2"
              placeholder="Email"
              value={loginData.email}
              onChange={(e) => setLoginData({ ...loginData, email: e.target.value })}
            />
            <input
              type="password"
              className="border p-2"
              placeholder="Password"
              value={loginData.password}
              onChange={(e) => setLoginData({ ...loginData, password: e.target.value })}
            />
            <button className="bg-blue-500 text-white p-2">Log In</button>
          </form>
        )}
        {tab === 'register' && (
          <form onSubmit={handleRegister} className="flex flex-col gap-2">
            <input
              type="email"
              className="border p-2"
              placeholder="Email"
              value={regData.email}
              onChange={(e) => setRegData({ ...regData, email: e.target.value })}
            />
            <input
              type="password"
              className="border p-2"
              placeholder="Password"
              value={regData.password}
              onChange={(e) => setRegData({ ...regData, password: e.target.value })}
            />
            <button className="bg-blue-500 text-white p-2">Register</button>
          </form>
        )}
      </div>
    );
  }

  return (
    <div className="max-w-xl mx-auto mt-10">
      <button className="mb-4 text-blue-600 underline" onClick={logout}>
        Logout
      </button>
      <ul className="bg-white shadow divide-y">
        {books.map((b) => (
          <li key={b.id} className="p-4 flex justify-between">
            <span>{b.title}</span>
            <span>{b.quantity}</span>
          </li>
        ))}
      </ul>
    </div>
  );
}

ReactDOM.render(<App />, document.getElementById('root'));
