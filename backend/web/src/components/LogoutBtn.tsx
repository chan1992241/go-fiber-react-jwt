import { useDispatch } from "react-redux";
import { useLogoutMutation } from "../app/api/authApiSlice";
import { setCredentials } from "../features/auth/authSlice";
import { useNavigate } from "react-router";

const LogoutBtn = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [logout, { isLoading }] = useLogoutMutation();
    const handleLogout = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const response = await logout("").unwrap();
            dispatch(setCredentials({ token: "", user: "" }));
            navigate("/");
        } catch (e: any) {
            console.error(e);
        }
    };
    return (
        <form onSubmit={handleLogout}>
            <button type="submit">Logout</button>
        </form>
    );
};

export default LogoutBtn;
