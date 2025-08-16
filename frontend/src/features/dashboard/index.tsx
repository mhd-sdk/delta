import { Header } from "@/components/layout/header";
import { Main } from "@/components/layout/main";
import { ProfileDropdown } from "@/components/profile-dropdown";
import { Search } from "@/components/search";
import { ThemeSwitch } from "@/components/theme-switch";
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
} from "@/components/ui/context-menu";
import { useEffect, useState } from "react";
import { Grid } from "./components/grid";
import { Panel, PanelType } from "./panel";
import { Range, Unit } from "./timerange";

const LOCAL_STORAGE_KEY = "dashboard_panels";

export default function Dashboard() {
  // Charger les panels depuis localStorage ou fallback à l'état par défaut
  const [panels, setPanels] = useState<Panel[]>(() => {
    if (typeof window !== "undefined") {
      const stored = localStorage.getItem(LOCAL_STORAGE_KEY);
      if (stored) {
        try {
          return JSON.parse(stored) as Panel[];
        } catch (e) {
          console.error("Erreur en lisant les panels depuis localStorage", e);
        }
      }
    }
    return [
      {
        data: {
          type: PanelType.Chart,
          config: {
            range: Range.oneDay,
            ticker: "nvda",
            timeframe: { n: 1, unit: Unit.hour },
          },
        },
        height: 5,
        width: 3,
        id: "1",
        x: 0,
        y: 0,
      },
    ];
  });

  const handleNewPanel = () => {
    const newPanel: Panel = {
      data: {
        type: PanelType.Chart,
        config: {
          range: Range.oneDay,
          ticker: "nvda",
          timeframe: { n: 1, unit: Unit.hour },
        },
      },
      height: 5,
      width: 3,
      id: Date.now().toString(), // Utiliser un timestamp comme ID unique
      x: 0,
      y: panels.length * 6, // Positionner le nouveau panel en dessous des précédents
    };
    setPanels([...panels, newPanel]);
  };

  // Sauvegarder les panels dans localStorage à chaque changement
  useEffect(() => {
    try {
      localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(panels));
    } catch (e) {
      console.error("Erreur en sauvegardant les panels dans localStorage", e);
    }
  }, [panels]);

  return (
    <>
      {/* ===== Top Heading ===== */}
      <Header>
        <div className="ml-auto flex items-center space-x-4">
          <Search />
          <ThemeSwitch />
          <ProfileDropdown />
        </div>
      </Header>

      {/* ===== Main ===== */}
      <ContextMenu>
        <ContextMenuTrigger>
          <Main className="h-screen">
            <Grid panels={panels} onChange={setPanels} />
          </Main>
        </ContextMenuTrigger>
        <ContextMenuContent className="w-64">
          <ContextMenuItem onClick={handleNewPanel}>New Panel</ContextMenuItem>
        </ContextMenuContent>
      </ContextMenu>
    </>
  );
}
