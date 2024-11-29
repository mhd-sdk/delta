import { GaugeChart, GaugeTypes } from '@carbon/charts-react';
import { Slider } from '@carbon/react';
import { css } from '@emotion/css';
import { useMarketMock } from '../../hooks/useMarketMock';
import { Graph } from '../Graph/Graph';

function getDateMinusSeconds(seconds: number): Date {
  const currentDate = new Date();
  const newDate = new Date(currentDate.getTime() - seconds * 1000);
  return newDate;
}

export const MarketMock = (): JSX.Element => {
  const { settings, tape } = useMarketMock();
  const { setAverageBuy, setAverageSell, setBuyFrequency, setSellFrequency, averageBuy, averageSell, buyFrequency, sellFrequency } = settings;
  const last30SecondsTape = tape.filter((t) => t.time > getDateMinusSeconds(5));
  const buyVolume = last30SecondsTape.filter((t) => t.side === 'buy').reduce((acc, curr) => acc + curr.size, 0);
  const sellVolume = last30SecondsTape.filter((t) => t.side === 'sell').reduce((acc, curr) => acc + curr.size, 0);
  const percentBuy = (buyVolume / (buyVolume + sellVolume)) * 100;

  return (
    <>
      <Graph />
      <div className={styles.container}>
        <GaugeChart
          data={[
            {
              group: 'value',
              value: percentBuy,
              min: -100,
              max: 100,
            },
            // {
            //   group: 'delta',
            //   value: -13.37,
            // },
          ]}
          options={{
            title: 'Delta',
            resizable: false,
            animations: true,
            height: '250px',
            gauge: {
              status: 'SUCCESS',
              numberFormatter: (value) => {
                switch (true) {
                  case value > 60:
                    return 'High buy volume';
                  case value < 40:
                    return 'High sell volume';
                }
                return 'Balanced';
              },
              type: GaugeTypes.SEMI,
              showPercentageSymbol: false,
              alignment: 'center',
            },
          }}
        />
        <div className={styles.flex}>
          <Slider
            labelText="Aggressive buyer (+-10)"
            value={averageBuy}
            min={0}
            max={100}
            stepMultiplier={50}
            onChange={(val) => setAverageBuy(val.value)}
            step={1}
            formatLabel={(val) => {
              if (val < 25) {
                return 'Low';
              } else if (val > 75) {
                return 'High';
              }
              return 'Medium';
            }}
          />
          <Slider
            labelText="Buy frequency (orders/ms)"
            value={buyFrequency}
            min={0}
            max={5000}
            stepMultiplier={50}
            onChange={(val) => setBuyFrequency(val.value)}
            step={1}
          />
        </div>
        <div className={styles.flex}>
          <Slider
            labelText="Aggressive seller (+-10)"
            value={averageSell}
            min={0}
            max={100}
            stepMultiplier={50}
            onChange={(val) => setAverageSell(val.value)}
            step={1}
          />
          <Slider
            labelText="Sell frequency (orders/ms)"
            value={sellFrequency}
            min={0}
            max={5000}
            stepMultiplier={50}
            onChange={(val) => setSellFrequency(val.value)}
            step={1}
          />
        </div>
      </div>
    </>
  );
};

const styles = {
  container: css`
    display: flex;
    flex-direction: column;
    gap: 16px;
    padding: 16px;
    align-items: center;
    justify-content: center;
  `,
  flex: css`
    display: flex;
    gap: 16px;
    align-items: center;
  `,
};
