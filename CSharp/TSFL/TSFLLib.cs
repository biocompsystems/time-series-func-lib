using System;

namespace TSFL
{
    public static class TSFLLib
    {
      /// <summary>
      /// Computes Simple Moving Average of specified period on full array passed
      /// </summary>
      /// <param name="data">The array of doubles to compute upon</param>
      /// <param name="period">The length of the moving average</param>
      /// <returns></returns>
      public static double[] SMA(double[] data, int period)
      {
        //merely convert to extended form, call and return the result
        return SMA(data, period, 0, data.Length);
      }

      /// <summary>
      /// Computes Simple Moving Average of specified period between start and end row
      /// </summary>
      /// <param name="data">The array of doubles to compute upon</param>
      /// <param name="period">The length of the moving average</param>
      /// <param name="startRow">The starting row of the computation (zero based)</param>
      /// <param name="endRow">The ending row of the computation (zero based)</param>
      /// <returns></returns>
      public static double[] SMA(double[] data, int period, int startRow, int endRow)
      {
        //Do checks
        if (data == null)
        {
          throw new ArgumentException("data cannot be null");
        }
        if (period < 0)
        {
          throw new ArgumentException("period cannot be negative");
        }
        if (startRow < 0)
        {
          throw new ArgumentException("startRow cannot be negative");
        }
        if (endRow < 0)
        {
          throw new ArgumentException("endRow cannot be negative");
        }
        if (endRow < startRow)
        {
          throw new ArgumentException("endRow must be greater or equal to startRow");
        }

        //allocate returned result the same as the incoming array
        var returnData = new double[data.Length];
        //if the period is greater than the length of the data, just return array with zeros
        if (period > data.Length) return returnData;

        //Figure out where to start computing.  It isn't just as simple as using start and end row
        //if the start row is less than the period, use the period, else the start row
        int startPoint = startRow < period ? period : startRow;
        
        //if the endrow is bigger than or equal to the data length, use the data length
        //else use the endrow + 1 as we are going one less than the endpoint: i < endPoint
        int endPoint;
        if (endRow >= data.Length)
        {
          endPoint = data.Length;
        }
        else
        {
          endPoint = endRow + 1;
        }

        //Walk through the data from start to end
        for (int i = startPoint - 1; i < endPoint; i++)
        {
          //computing a moving window of size period
          double value = 0;
          for (int j = i; j > i - period; j += -1)
          {
            value += data[j];
          }
          returnData[i] = value / period;
        }
        return returnData;
      }
    }
}

