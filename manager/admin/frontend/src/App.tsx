import CssBaseline from "@mui/material/CssBaseline";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import useMediaQuery from "@mui/material/useMediaQuery";

import HomeView from "./views/Home";
import { RecoilRoot, useRecoilValue } from "recoil";
import Loading from "./components/Loading";
import { loadingState } from "./state/atoms";

function App() {
  const isDarkMode = useMediaQuery("(prefers-color-scheme: dark)");
  const isLoading = useRecoilValue(loadingState);
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
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <HomeView />
      {isLoading && <Loading />}
    </ThemeProvider>
  );
}

export default App;
