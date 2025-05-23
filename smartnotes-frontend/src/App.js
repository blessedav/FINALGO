import React, { useState, useEffect } from "react";

const API = "http://localhost:3001/api";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token") || "");
  const [books, setBooks] = useState([]);
  const [form, setForm] = useState({ title: "", author: "", description: "", tags: "" });
  const [auth, setAuth] = useState({ email: "", password: "", username: "" });
  const [isLogin, setIsLogin] = useState(true);
  const [viewBook, setViewBook] = useState(null);
  const [editingBook, setEditingBook] = useState(null);

  // Получить книги
  useEffect(() => {
    if (token) {
      fetch(`${API}/books`, {
        headers: { Authorization: `Bearer ${token}` }
      })
        .then(async res => {
          if (!res.ok) {
            const err = await res.json().catch(() => ({}));
            console.error("Ошибка при получении книг:", err);
            alert(err.error || "Ошибка при получении книг");
            return { books: [] };
          }
          return res.json();
        })
        .then(data => setBooks(data.books || data || []));
    }
  }, [token]);

  // Аутентификация
  const handleAuth = async (e) => {
    e.preventDefault();
    const url = isLogin ? `${API}/auth/login` : `${API}/auth/register`;
    const body = isLogin
      ? { email: auth.email, password: auth.password }
      : { username: auth.username, email: auth.email, password: auth.password };
    try {
      const res = await fetch(url, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });
      const data = await res.json().catch(() => ({}));
      console.log("Ответ сервера:", data);
      if (data.token) {
        setToken(data.token);
        localStorage.setItem("token", data.token);
      } else {
        alert(data.error || "Auth error");
      }
    } catch (err) {
      console.error("Ошибка сети при аутентификации:", err);
      alert("Ошибка сети при аутентификации");
    }
  };

  // Создать книгу
  const handleCreate = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(`${API}/books`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          ...form,
          tags: form.tags.split(",").map((t) => t.trim()),
        }),
      });
      const data = await res.json().catch(() => ({}));
      if (res.ok) {
        setForm({ title: "", author: "", description: "", tags: "" });
        // обновить список книг
        fetch(`${API}/books`, {
          headers: { Authorization: `Bearer ${token}` },
        })
          .then(async res => {
            if (!res.ok) {
              const err = await res.json().catch(() => ({}));
              console.error("Ошибка при получении книг:", err);
              alert(err.error || "Ошибка при получении книг");
              return { books: [] };
            }
            return res.json();
          })
          .then(data => setBooks(data.books || data || []));
      } else {
        console.error("Ошибка при создании книги:", data);
        alert(data.error || "Ошибка при создании книги");
      }
    } catch (err) {
      console.error("Ошибка сети при создании книги:", err);
      alert("Ошибка сети при создании книги");
    }
  };

  // Удалить книгу
  const handleDelete = async (id) => {
    try {
      const res = await fetch(`${API}/books/${id}`, {
        method: "DELETE",
        headers: { Authorization: `Bearer ${token}` },
      });
      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        console.error("Ошибка при удалении книги:", data);
        alert(data.error || "Ошибка при удалении книги");
        return;
      }
      setBooks(books.filter((b) => b.id !== id && b._id !== id));
    } catch (err) {
      console.error("Ошибка сети при удалении книги:", err);
      alert("Ошибка сети при удалении книги");
    }
  };

  const handleView = async (id) => {
    try {
      const res = await fetch(`${API}/books/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      const data = await res.json().catch(() => ({}));
      if (!res.ok) {
        console.error("Ошибка при получении книги:", data);
        alert(data.error || "Ошибка при получении книги");
        return;
      }
      setViewBook(data);
    } catch (err) {
      console.error("Ошибка сети при получении книги:", err);
      alert("Ошибка сети при получении книги");
    }
  };

  const startEdit = (book) => {
    setEditingBook(book);
    setForm({
      title: book.title,
      author: book.author,
      description: book.description,
      tags: (book.tags || []).join(", "),
    });
  };

  const handleUpdate = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(`${API}/books/${editingBook.id || editingBook._id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          ...form,
          tags: form.tags.split(",").map((t) => t.trim()),
        }),
      });
      const data = await res.json().catch(() => ({}));
      if (res.ok) {
        setEditingBook(null);
        setForm({ title: "", author: "", description: "", tags: "" });
        // обновить список книг
        fetch(`${API}/books`, {
          headers: { Authorization: `Bearer ${token}` },
        })
          .then(async res => {
            if (!res.ok) {
              const err = await res.json().catch(() => ({}));
              console.error("Ошибка при получении книг:", err);
              alert(err.error || "Ошибка при получении книг");
              return { books: [] };
            }
            return res.json();
          })
          .then(data => setBooks(data.books || data || []));
      } else {
        console.error("Ошибка при обновлении книги:", data);
        alert(data.error || "Ошибка при обновлении книги");
      }
    } catch (err) {
      console.error("Ошибка сети при обновлении книги:", err);
      alert("Ошибка сети при обновлении книги");
    }
  };

  if (!token) {
    return (
      <div style={{ maxWidth: 400, margin: "40px auto" }}>
        <h2>{isLogin ? "Вход" : "Регистрация"}</h2>
        <form onSubmit={handleAuth}>
          {!isLogin && (
            <input
              placeholder="Username"
              value={auth.username}
              onChange={e => setAuth({ ...auth, username: e.target.value })}
              required
            />
          )}
          <input
            placeholder="Email"
            type="email"
            value={auth.email}
            onChange={e => setAuth({ ...auth, email: e.target.value })}
            required
          />
          <input
            placeholder="Password"
            type="password"
            value={auth.password}
            onChange={e => setAuth({ ...auth, password: e.target.value })}
            required
          />
          <button type="submit">{isLogin ? "Войти" : "Зарегистрироваться"}</button>
        </form>
        <button onClick={() => setIsLogin(!isLogin)}>
          {isLogin ? "Нет аккаунта? Регистрация" : "Уже есть аккаунт? Войти"}
        </button>
      </div>
    );
  }

  return (
    <div style={{ maxWidth: 600, margin: "40px auto" }}>
      <h2>Книги</h2>
      <form onSubmit={editingBook ? handleUpdate : handleCreate}>
        <input
          placeholder="Название"
          value={form.title}
          onChange={e => setForm({ ...form, title: e.target.value })}
          required
        />
        <input
          placeholder="Автор"
          value={form.author}
          onChange={e => setForm({ ...form, author: e.target.value })}
        />
        <input
          placeholder="Описание"
          value={form.description}
          onChange={e => setForm({ ...form, description: e.target.value })}
        />
        <input
          placeholder="Теги (через запятую)"
          value={form.tags}
          onChange={e => setForm({ ...form, tags: e.target.value })}
        />
        <button type="submit">{editingBook ? "Сохранить" : "Добавить книгу"}</button>
        {editingBook && (
          <button type="button" onClick={() => { setEditingBook(null); setForm({ title: "", author: "", description: "", tags: "" }); }}>
            Отмена
          </button>
        )}
      </form>
      <ul>
        {books.map((b) => (
          <li key={b.id}>
            <b>{b.title}</b> — {b.author} <br />
            <i>{b.description}</i> <br />
            Теги: {b.tags && b.tags.join(", ")}
            <button onClick={() => handleDelete(b.id)} style={{ marginLeft: 10 }}>
              Удалить
            </button>
            <button onClick={() => startEdit(b)} style={{ marginLeft: 10 }}>
              Редактировать
            </button>
            <button onClick={() => handleView(b.id)} style={{ marginLeft: 10 }}>
              Посмотреть
            </button>
          </li>
        ))}
      </ul>
      <button onClick={() => { setToken(""); localStorage.removeItem("token"); }}>
        Выйти
      </button>
      {viewBook && (
        <div style={{ border: "1px solid #ccc", padding: 10, margin: 10 }}>
          <h3>{viewBook.title}</h3>
          <p>Автор: {viewBook.author}</p>
          <p>Описание: {viewBook.description}</p>
          <p>Теги: {viewBook.tags && viewBook.tags.join(", ")}</p>
          <button onClick={() => setViewBook(null)}>Закрыть</button>
        </div>
      )}
    </div>
  );
}

export default App;
