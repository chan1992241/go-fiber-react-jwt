import { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router";
import { useRegisterMutation } from "../../app/api/authApiSlice";
import { Link } from "react-router-dom";

export default function Register() {
    const userRef = useRef<HTMLInputElement>(null);
    const errRef = useRef<HTMLInputElement>(null);
    const [user, setUser] = useState("");
    const [pwd, setPwd] = useState("");
    const [errMsg, setErrMsg] = useState("");
    const [register, { isLoading }] = useRegisterMutation();
    const navigate = useNavigate();
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            await register({
                username: user,
                password: pwd,
                role: "employee",
            }).unwrap();
            navigate("/");
        } catch (err: any) {
            if (!err?.originalStatus) {
                // isLoading: true until timeout occurs
                setErrMsg("No Server Response");
            } else if (err.originalStatus === 400) {
                setErrMsg("Missing Username or Password");
            } else if (err.originalStatus === 500) {
                setErrMsg("Server Error");
            } else if (err.originalStatus == 409) {
                setErrMsg("Username already exists");
            } else {
                setErrMsg("Register Failed");
            }
        }
    };
    useEffect(() => {
        userRef.current && userRef.current.focus();
    }, []);
    return (
        <>
            <section className="login">
                <p
                    ref={errRef}
                    className={errMsg ? "errmsg" : "offscreen"}
                    aria-live="assertive"
                >
                    {errMsg}
                </p>

                <h1>Register Employee</h1>

                <form onSubmit={handleSubmit}>
                    <label htmlFor="username">Username:</label>
                    <input
                        type="text"
                        id="username"
                        ref={userRef}
                        value={user}
                        autoComplete="true"
                        required
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                            setUser(e.target.value)
                        }
                    />

                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        required
                        value={pwd}
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                            setPwd(e.target.value);
                        }}
                    />
                    <button>Register</button>
                </form>
                <Link to="/">Back</Link>
            </section>
        </>
    );
}
