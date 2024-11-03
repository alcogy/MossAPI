import CssBaseline from "@mui/material/CssBaseline";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import useMediaQuery from "@mui/material/useMediaQuery";

import HomeView from "./views/Home";
import { RecoilRoot } from "recoil";

function App() {
  const isDarkMode = useMediaQuery("(prefers-color-scheme: dark)");

  const theme = createTheme({
    palette: {
      mode: isDarkMode ? "dark" : "light",
      primary: {
        main: "#345247",
        light: "#009688",
      },
      secondary: {
        main: "#cccccc",
      },
    },

    components: {
      MuiAppBar: {
        styleOverrides: {
          colorPrimary: {
            backgroundColor: "#345247",
          },
        },
      },
      MuiButton: {
        styleOverrides: {
          containedPrimary: {
            backgroundColor: "#347267",
          },
        },
      },
    },
  });

  return (
    <RecoilRoot>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <HomeView />
      </ThemeProvider>
    </RecoilRoot>
  );
}

export default App;
