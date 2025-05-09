import React, { useState } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import Brightness4Icon from "@mui/icons-material/Brightness4";
import Brightness7Icon from "@mui/icons-material/Brightness7";
import SettingsIcon from "@mui/icons-material/Settings";
import Home from "./pages/HomePage";
import MediaItemPage from "./pages/MediaItemPage";
import PageNotFound from "./pages/NotFoundPage";
import SettingsPage from "./pages/SettingsPage";
import { Save, SettingsApplications } from "@mui/icons-material";
import LogsPage from "./pages/LogsPage";
import SavedSetsPage from "./pages/SavedSetsPage";

type AppProps = {
	darkMode: boolean;
	setDarkMode: React.Dispatch<React.SetStateAction<boolean>>;
};

function App({ darkMode, setDarkMode }: AppProps) {
	const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
	const navigate = useNavigate(); // Move useNavigate to the top level of the component

	const handleMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
		setAnchorEl(event.currentTarget);
	};

	const handleMenuClose = () => {
		setAnchorEl(null);
	};

	const toggleDarkMode = () => {
		setDarkMode((prevMode) => !prevMode);
		handleMenuClose();
	};

	return (
		<>
			<AppBar position="fixed">
				<Toolbar>
					{/* Logo */}
					<img
						src="/logo.png"
						alt="Logo"
						style={{ height: 40, cursor: "pointer" }}
						onClick={() => navigate("/")} // Use navigate function here
					/>

					{/* Spacer to push the settings button to the right */}
					<div style={{ flexGrow: 1 }}></div>

					{/* Settings Button */}
					<IconButton
						edge="end"
						color="inherit"
						aria-label="settings"
						onClick={handleMenuOpen}
					>
						<SettingsIcon />
					</IconButton>

					{/* Settings Menu */}
					<Menu
						anchorEl={anchorEl}
						open={Boolean(anchorEl)}
						onClose={handleMenuClose}
						slotProps={{
							transition: {
								timeout: 0,
							},
						}}
					>
						<MenuItem onClick={toggleDarkMode}>
							{darkMode ? (
								<>
									<Brightness7Icon sx={{ marginRight: 1 }} />{" "}
									Light Mode
								</>
							) : (
								<>
									<Brightness4Icon sx={{ marginRight: 1 }} />{" "}
									Dark Mode
								</>
							)}
						</MenuItem>
						<MenuItem
							onClick={() => {
								handleMenuClose();
								navigate("/settings");
							}}
						>
							<>
								<SettingsApplications sx={{ marginRight: 1 }} />{" "}
								Settings Page
							</>
						</MenuItem>
						<MenuItem
							onClick={() => {
								handleMenuClose();
								navigate("/saved_sets");
							}}
						>
							<>
								<Save sx={{ marginRight: 1 }} /> Saved Sets
							</>
						</MenuItem>
					</Menu>
				</Toolbar>
			</AppBar>

			{/* Routes */}
			<div style={{ paddingTop: 64 }}>
				<Routes>
					<Route path="/" element={<Home />} />
					<Route
						path="/media/:ratingKey/:title"
						element={<MediaItemPage />}
					/>
					<Route path="/settings" element={<SettingsPage />} />
					<Route path="/logs" element={<LogsPage />} />
					<Route path="/saved_sets" element={<SavedSetsPage />} />

					{/* 404 Page Not Found */}
					<Route path="*" element={<PageNotFound />} />
				</Routes>
			</div>
		</>
	);
}

export default App;
