using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace TSFLUnitTestProject
{
  [TestClass]
  public class TSFLUnitTest
  {
    [TestMethod][TestCategory("MA")]
    public void SMA_Simple_Test()
    {
      //Send a small array of same lenght as period
      var array = new double[] {1.0, 2.0, 3.0};
      var result = TSFL.TSFLLib.SMA(array, 3);
      //result should come back with a 2 in the last position (1+2+3)/3
      Assert.AreEqual(2.0, result[2]);

      //Send a small array of same lenght as period
      array = new double[] { 1.0, 2.0, 3.0, 4.0, 5.0 };
      result = TSFL.TSFLLib.SMA(array, 3);
      //result should come back with a 4 in the last position (3+4+5)/3
      Assert.AreEqual(4.0, result[4]);

      //Call with additional start and end parameters set to beginning and end of array
      result = TSFL.TSFLLib.SMA(array, 3, 0, array.Length);
      Assert.AreEqual(4.0, result[4]);

      //Call with intermediate range to compute
      result = TSFL.TSFLLib.SMA(array, 3, 1, 3);
      Assert.AreEqual(4.0, result[4]);
    }
  }
}
