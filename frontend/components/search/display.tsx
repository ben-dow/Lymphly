import {
  Box,
  Button,
  Divider,
  LoadingOverlay,
  Select,
  Tabs,
  TextInput,
} from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { Position } from "geojson";
import { LngLatLike } from "maplibre-gl";
import Radar from "radar-sdk-js";
import RadarMap from "radar-sdk-js/dist/ui/RadarMap";
import { JSX, useEffect, useState } from "react";
import {
  LimitedPracticePracticeListI as LimitedPracticesListI,
  PracticeI,
  PracticeListI,
  ProviderListI,
} from "../../model/practice";

export function DataDisplay() {
  const [selectedPractice, setSelectedPractice] = useState("");
  const [practices, setPractices] = useState<LimitedPracticesListI>({
    practices: [],
  });
  const [mapCfg, setMapCfg] = useState<MapConfiguration>();
  const [visible, { open, close }] = useDisclosure(false);

  return (
    <Box pos={"relative"} className="flex flex-col w-full md:w-7xl">
      <LoadingOverlay
        visible={visible}
        zIndex={1000}
        overlayProps={{ radius: "md", blur: 1 }}
      />
      <Box className="w-full flex justify-center rounded-tl-xl bg-white">
        <Tabs
          defaultValue={"browse"}
          className="w-full"
          orientation="vertical"
          keepMounted={false}
        >
          <Tabs.List className="bg-cyan-600 rounded-t-xl md:rounded-tl-xl md:rounded-tr-none">
            <Tabs.Tab value="browse">
              <h1 className="font-sans text-white font-medium">Browse</h1>
            </Tabs.Tab>
            <Tabs.Tab value="current">
              <h1 className="font-sans text-white font-medium">
                Current Location
              </h1>
            </Tabs.Tab>
            <Tabs.Tab value="addr">
              <h1 className="font-sans text-white font-medium">Address</h1>
            </Tabs.Tab>
            <Tabs.Tab value="state">
              <h1 className="font-sans text-white font-medium">State</h1>
            </Tabs.Tab>
          </Tabs.List>
          <Tabs.Panel
            value="browse"
            className="flex justify-center w-full h-full"
          >
            <Browse
              loadingOn={open}
              loadingOff={close}
              setMapCfg={setMapCfg}
              setPractices={setPractices}
            />
          </Tabs.Panel>
          <Tabs.Panel
            value="current"
            className="flex justify-center w-full h-full"
          >
            <SearchByLocation
              loadingOn={open}
              loadingOff={close}
              setMapCfg={setMapCfg}
              setPractices={setPractices}
            />
          </Tabs.Panel>
          <Tabs.Panel value="addr" className="h-full w-full">
            <SearchByAddress
              loadingOn={open}
              loadingOff={close}
              setMapCfg={setMapCfg}
              setPractices={setPractices}
            />
          </Tabs.Panel>
          <Tabs.Panel value="state" className="h-full w-full">
            <SearchByState
              loadingOn={open}
              loadingOff={close}
              setMapCfg={setMapCfg}
              setPractices={setPractices}
            />
          </Tabs.Panel>
        </Tabs>
      </Box>
      <Box className="flex flex-col md:flex-row w-full justify-center border-t">
        <Tabs defaultValue={"Map"} className="w-full md:w-7/8 max-w-5xl">
          <Tabs.List className="bg-cyan-700  md:rounded-tr-none">
            <Tabs.Tab value="Map">
              <h1 className="font-sans text-white font-medium">Map</h1>
            </Tabs.Tab>
            <Tabs.Tab value="Table">
              <h1 className="font-sans text-white font-medium">List</h1>
            </Tabs.Tab>
          </Tabs.List>
          <Tabs.Panel value="Map" className="flex justify-center w-full h-full">
            <Map
              mapConfiguration={mapCfg}
              setSelectedPractice={setSelectedPractice}
              practiceList={practices}
            />
          </Tabs.Panel>
          <Tabs.Panel value="Table" className="h-full w-full">
            <PracticeTable
              practiceList={practices}
              updatedSelected={setSelectedPractice}
            />
          </Tabs.Panel>
        </Tabs>
        <Box className="md:h-full md:w-xs min-h-25 w-full  bg-gray-100 ">
          <Selected practiceId={selectedPractice} />
        </Box>
      </Box>
    </Box>
  );
}

interface PracticeUpdaterI {
  setPractices: (props: LimitedPracticesListI) => void;
  setMapCfg: (mapCfg: MapConfiguration) => void;
  loadingOn: () => void;
  loadingOff: () => void;
}

function Browse(props: PracticeUpdaterI) {
  useEffect(() => {
    props.loadingOn();
    fetch("/api/public/v1/practices/all")
      .then((r) => r.json())
      .then((j) => {
        const pl: PracticeListI = j;
        props.setPractices(pl);
        props.setMapCfg({ RadiusFeature: false });
      })
      .finally(props.loadingOff);
  }, []);

  return (
    <Box className="w-full h-full flex justify-center items-center">
      <Box className="text-xl font-sans font-semibold">
        Now Displaying All Practices
      </Box>
    </Box>
  );
}

function SearchByLocation(props: PracticeUpdaterI) {
  const [radius, setRadius] = useState(25);

  useEffect(() => {
    navigator.geolocation.getCurrentPosition((pos) => {
      let lat = pos.coords.latitude;
      let long = pos.coords.longitude;
      props.loadingOn();
      fetch(
        `/api/public/v1/practices/locate/proximity?lat=${lat}&long=${long}&radius=${radius}`,
      )
        .then((res) => res.json())
        .then((res) => {
          props.setPractices(res);
          props.setMapCfg({
            RadiusFeature: true,
            RadiusOrigin: [long, lat],
            Radius: radius,
          });
        })
        .finally(() => {
          props.loadingOff();
        });
    });
  }, [radius]);

  return (
    <div>
      <Box className="flex justify-center h-full flex-col p-2">
        <Box className="bg-white rounded  p-4 flex flex-col gap-2">
          <Box className="flex flex-col sm:flex-row gap-5 ">
            <Box className="font-sans text-xl font-medium text-sky-950 flex flex-col justify-center">
              Search Radius:
            </Box>
            <Button
              onClick={() => {
                setRadius(25);
              }}
            >
              25 Miles{" "}
            </Button>
            <Button
              onClick={() => {
                setRadius(50);
              }}
            >
              50 Miles{" "}
            </Button>
            <Button
              onClick={() => {
                setRadius(100);
              }}
            >
              100 Miles{" "}
            </Button>
          </Box>
          <Box className="text-center">Current: {radius} Miles</Box>
        </Box>
      </Box>
    </div>
  );
}

function SearchByAddress(props: PracticeUpdaterI) {
  const [addr, setAddr] = useState("");
  const [radius, setRadius] = useState(25);
  const [firstRequestMade, setFirstRequestMade] = useState(false);

  const update = () => {
    props.loadingOn();
    fetch(
      `/api/public/v1/practices/locate/proximity?addr=${btoa(addr)}&radius=${radius}`,
    )
      .then((res) => res.json())
      .then((res) => {
        props.setPractices(res);
        props.setMapCfg({
          RadiusFeature: true,
          RadiusOrigin: [res["originLongitude"], res["originLatitude"]],
          Radius: radius,
        });
      })
      .finally(() => {
        props.loadingOff();
      });
  };

  useEffect(() => {
    if (addr != "" && firstRequestMade) {
      update();
    }
  }, [radius]);

  return (
    <Box className="w-full h-full p-2 flex flex-col gap-2">
      <Box className="flex justify-center gap-2">
        <Box className="w-100">
          <TextInput
            placeholder="Enter Address or City Name"
            value={addr}
            onChange={(e) => {
              setAddr(e.currentTarget.value);
            }}
          />
        </Box>
        <Box>
          <Button
            onClick={() => {
              setFirstRequestMade(true);
              update();
            }}
          >
            Search
          </Button>
        </Box>
      </Box>

      <Box className="flex flex-col items-center gap-2 justify-center">
        <Box className="flex gap-2">
          <Button
            onClick={() => {
              setRadius(25);
            }}
          >
            25 Miles{" "}
          </Button>
          <Button
            onClick={() => {
              setRadius(50);
            }}
          >
            50 Miles{" "}
          </Button>
          <Button
            onClick={() => {
              setRadius(100);
            }}
          >
            100 Miles{" "}
          </Button>
        </Box>
        <Box className="text-center font-sans">Current: {radius} Miles</Box>
      </Box>
    </Box>
  );
}

function SearchByState(props: PracticeUpdaterI) {
  var usStates = {
    ALABAMA: "AL",
    ALASKA: "AK",
    "AMERICAN SAMOA": "AS",
    ARIZONA: "AZ",
    ARKANSAS: "AR",
    CALIFORNIA: "CA",
    COLORADO: "CO",
    CONNECTICUT: "CT",
    DELAWARE: "DE",
    "DISTRICT OF COLUMBIA": "DC",
    FLORIDA: "FL",
    GEORGIA: "GA",
    GUAM: "GU",
    HAWAII: "HI",
    IDAHO: "ID",
    ILLINOIS: "IL",
    INDIANA: "IN",
    IOWA: "IA",
    KANSAS: "KS",
    KENTUCKY: "KY",
    LOUISIANA: "LA",
    MAINE: "ME",
    "MARSHALL ISLANDS": "MH",
    MARYLAND: "MD",
    MASSACHUSETTS: "MA",
    MICHIGAN: "MI",
    MINNESOTA: "MN",
    MISSISSIPPI: "MS",
    MISSOURI: "MO",
    MONTANA: "MT",
    NEBRASKA: "NE",
    NEVADA: "NV",
    "NEW HAMPSHIRE": "NH",
    "NEW JERSEY": "NJ",
    "NEW MEXICO": "NM",
    "NEW YORK": "NY",
    "NORTH CAROLINA": "NC",
    "NORTH DAKOTA": "ND",
    "NORTHERN MARIANA ISLANDS": "MP",
    OHIO: "OH",
    OKLAHOMA: "OK",
    OREGON: "OR",
    PALAU: "PW",
    PENNSYLVANIA: "PA",
    "PUERTO RICO": "PR",
    "RHODE ISLAND": "RI",
    "SOUTH CAROLINA": "SC",
    "SOUTH DAKOTA": "SD",
    TENNESSEE: "TN",
    TEXAS: "TX",
    UTAH: "UT",
    VERMONT: "VT",
    "VIRGIN ISLANDS": "VI",
    VIRGINIA: "VA",
    WASHINGTON: "WA",
    "WEST VIRGINIA": "WV",
    WISCONSIN: "WI",
    WYOMING: "WY",
  };
  return (
    <Box className="flex justify-center font-sans text-xl items-center gap-2 w-full h-full">
      <Box>Select State: </Box>
      <Select
        data={Object.keys(usStates)}
        onChange={(res) => {
          props.loadingOn();
          fetch(
            `/api/public/v1/practices/locate/state/` +
            usStates[res.valueOf()],
          )
            .then((res) => res.json())
            .then((res) => {
              props.setPractices(res);
              props.setMapCfg({
                RadiusFeature: false,
              });
            })
            .finally(props.loadingOff);
        }}
      />
    </Box>
  );
}

interface SelectedProps {
  practiceId: string;
}

function Selected(props: SelectedProps) {
  const [practice, setPractice] = useState<PracticeI>();
  const [providers, setProviders] = useState<ProviderListI>({ providers: [] });
  const [practiceLoading, practiceLoader] = useDisclosure(false);
  const [providerLoading, providerLoader] = useDisclosure(false);

  useEffect(() => {
    if (props.practiceId != "") {
      practiceLoader.open();
      fetch("/api/public/v1/practice/" + props.practiceId)
        .then((r) => r.json())
        .then((pr) => {
          setPractice(pr);
        })
        .finally(practiceLoader.close);

      providerLoader.open();
      fetch(
        "/api/public/v1/practice/" + props.practiceId + "/providers",
      )
        .then((r) => r.json())
        .then((pr) => {
          setProviders(pr);
        })
        .finally(providerLoader.close);
    }
  }, [props.practiceId]);

  if (practice === undefined) {
    return (
      <Box
        pos="relative"
        className="flex flex-col gap-5 border-l-2 border-black h-full"
      >
        <LoadingOverlay
          visible={practiceLoading && providerLoading}
          zIndex={1000}
          overlayProps={{ radius: "md", blur: 1 }}
        />
        <Box className="text-center font-sans text-2xl text-white font-medium bg-cyan-700 h-9 border-b-2">
          Selected Practice
        </Box>

        <Box className="text-center font-sans text-2xl text-black font-medium2">
          No Practice Selected
        </Box>
      </Box>
    );
  } else {
    let rows: JSX.Element[] = [];
    rows = providers.providers.map((r, idx) => {
      return (
        <Box className="">
          <Box
            key={idx}
            className="text-lg rounded bg-white text-gray-900 font-sans"
          >
            {r.name}
          </Box>
          <Box className="font-semibold font-sans text-xs flex gap-1">
            <Box>Tags:</Box> {r.tags}
          </Box>
          <Divider />
        </Box>
      );
    });

    return (
      <Box
        pos="relative"
        className="flex flex-col border-l-2 border-black h-159"
      >
        <LoadingOverlay
          visible={practiceLoading && providerLoading}
          zIndex={1000}
          overlayProps={{ radius: "md", blur: 1 }}
        />
        <Box className="text-center font-sans text-2xl text-white font-medium bg-cyan-700 h-9 border-b-2">
          Selected Practice
        </Box>
        <Box className="h-full p-5 bg-white overflow-auto">
          <Box className="flex flex-col gap-1 justify-center text-center">
            <Box className="text-md font-semibold text-center">
              {practice.name}
            </Box>
            <Box className="text-sm text-center">{practice.fullAddress}</Box>
            <Box className="text-sm">
              <a
                target="_blank"
                rel="noopener noreferrer"
                className="underline"
                href={practice.website}
              >
                {practice.website}
              </a>
            </Box>
            <Box className="text-sm text-center">{practice.phone}</Box>
            <Box className="text-sm">{practice.tags}</Box>
          </Box>
          <Divider my="sm" labelPosition="center" label="Providers" />
          <Box pos="relative" className="flex flex-col gap-2">
            {rows}
          </Box>
        </Box>
      </Box>
    );
  }
}

interface PracticeTableProps {
  practiceList: LimitedPracticesListI;
  updatedSelected: (practiceId: string) => void;
}

export function PracticeTable(props: PracticeTableProps) {
  let rows: JSX.Element[] = [];

  rows = props.practiceList.practices.map((r, idx) => {
    return (
      <Box key={idx}>
        <Box
          onClick={() => {
            props.updatedSelected(r.practiceId);
          }}
          className="font-sans p-2 font-medium hover:bg-slate-100 outline m-1 hover:cursor-pointer rounded bg-white text-gray-900"
        >
          {r.name}
        </Box>
      </Box>
    );
  });

  return (
    <Box className="w-full h-75 md:h-150 md:rounded-bl-2xl bg-gray-100 p-5">
      <Box className="overflow-auto h-full">{rows}</Box>
    </Box>
  );
}

interface MapProps {
  mapConfiguration?: MapConfiguration;
  practiceList: LimitedPracticesListI;
  setSelectedPractice: (practiceId: string) => void;
}

interface MapConfiguration {
  RadiusOrigin?: LngLatLike;
  RadiusFeature?: boolean;
  Radius?: number;
}

export function Map(props: MapProps) {
  const [map, setMap] = useState<RadarMap>(undefined);

  useEffect(() => {
    if (map === undefined) {
      fetch("/radar_pub_key.txt")
        .then((r) => r.text())
        .then((text) => {
          Radar.initialize(text, { debug: false });
          const Map = Radar.ui.map({
            container: "map",
            zoom: 0,
          });
          setMap(Map);
        });
    }
  }, [map]);

  useEffect(() => {
    if (map != undefined) {
      if (props.practiceList != undefined) {
        map.clearMarkers();
        for (
          let index = 0;
          index < props.practiceList.practices.length;
          index++
        ) {
          const element = props.practiceList.practices[index];
          const marker = Radar.ui.marker({
            color: "red",
            scale: 0.5,
          });
          marker.on("click", () => {
            if (props.setSelectedPractice != undefined) {
              props.setSelectedPractice(element.practiceId);
            }
          });
          marker.setLngLat([element.longitude, element.lattitude]).addTo(map);
        }
        map.fitToMarkers();
      }

      map.clearFeatures();

      if (
        props.mapConfiguration != undefined &&
        props.mapConfiguration.RadiusFeature
      ) {
        const marker = Radar.ui.marker({
          color: "blue",
          scale: 0.75,
        });
        marker.setLngLat(props.mapConfiguration.RadiusOrigin).addTo(map);
        map.addPolygon(
          {
            type: "Feature",
            properties: {
              name: "radius",
            },
            geometry: {
              type: "Polygon",
              coordinates: ZoneCoords(
                props.mapConfiguration.RadiusOrigin,
                props.mapConfiguration.Radius,
                100,
              ),
            },
          },
          {
            paint: {
              "fill-color": "yellow",
              "fill-opacity": 0.1,
              "border-width": 1,
              "border-color": "red",
              "border-opacity": 0.3,
            },
          },
        );
        map.fitToFeatures();
      }

      map.redraw();
    }
  }, [props.practiceList, props.mapConfiguration, map]);

  return <div id="map" className="w-full h-75 md:h-150 md:rounded-b-2xl" />;
}

function ZoneCoords(
  lngLt: LngLatLike,
  radius: number,
  resolution: number,
): Position[][] {
  const long = lngLt[0];
  const lat = lngLt[1];

  const radiusKm = radius / 0.621371;
  const radiusLon =
    (1 / (111.319 * Math.cos(lat * (Math.PI / 180)))) * radiusKm;
  const radiusLat = (1 / 110.574) * radiusKm;

  const dTheta = (2 * Math.PI) / resolution;
  let theta = 0;

  let out: Position[] = [];

  for (var i = 0; i < resolution; i++) {
    out.push([
      long + radiusLon * Math.cos(theta),
      lat + radiusLat * Math.sin(theta),
    ]);
    theta += dTheta;
  }

  return [out];
}
