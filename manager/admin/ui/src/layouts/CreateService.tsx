import ModuleTitle from "../components/ModuleTitle";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import {
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  TextField,
} from "@mui/material";
import { useState } from "react";
import { API_GET_SERVICES, API_SERVICE_CREATE } from "../common/constants";
import { loadingState, serviceListState } from "../state/atoms";
import { useRecoilState } from "recoil";
import { Service } from "../state/models";

type BaseImage = "debian" | "centos" | "alpine" | "python";
interface ServiceForm {
  service: string;
  base: BaseImage;
  artifact: string;
  options: string;
  execute: string;
}

const initServiceForm: ServiceForm = {
  service: "",
  base: "debian",
  artifact: "",
  options: "",
  execute: "",
};

export default function CreateService() {
  const [form, setForm] = useState<ServiceForm>(initServiceForm);
  const [_serviceList, setServiceList] = useRecoilState(serviceListState);
  const [_isLoading, setIsLoading] = useRecoilState(loadingState);

  // Click create button handler.
  const onClickCreate = async () => {
    setIsLoading(true);
    try {
      const result = await fetch(API_SERVICE_CREATE, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(form),
      });
      const data = await result.json();
      if (data["message"] === "ok") {
        const response = await fetch(API_GET_SERVICES);
        const data = await response.json();
        setServiceList(data as Service[]);
        setForm(initServiceForm);
      }
    } catch (e) {
      console.error(e);
      alert("Sorry, you got error.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Box>
      <ModuleTitle label="Create Service" />
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: 3,
          maxWidth: "480px",
        }}
      >
        <TextField
          label="Service name"
          variant="outlined"
          value={form.service}
          onChange={(e) => setForm({ ...form, service: e.target.value })}
        />
        <FormControl sx={{ minWidth: "128px" }}>
          <InputLabel id="select-label">Base image</InputLabel>
          <Select
            label="Base image"
            labelId="select-label"
            value={form.base}
            onChange={(e) =>
              setForm({ ...form, base: e.target.value as BaseImage })
            }
          >
            {["debian", "centos", "alpine", "python"].map((v) => (
              <MenuItem value={v} key={v}>
                {v}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <TextField
          label="Root directory for artifact"
          variant="outlined"
          value={form.artifact}
          onChange={(e) => setForm({ ...form, artifact: e.target.value })}
        />
        <TextField
          label="Optional Dockerfile Commands."
          variant="outlined"
          multiline
          rows={6}
          value={form.options}
          onChange={(e) => setForm({ ...form, options: e.target.value })}
        />
        <TextField
          label="Execute command when start contaier."
          variant="outlined"
          value={form.execute}
          onChange={(e) => setForm({ ...form, execute: e.target.value })}
        />

        <Box sx={{ marginTop: 3, display: "flex", gap: 2 }}>
          <Button variant="contained" onClick={onClickCreate}>
            Create
          </Button>
          <Button variant="outlined" color="secondary" href="/">
            Cancel
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
