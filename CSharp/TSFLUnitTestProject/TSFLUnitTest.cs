using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace TSFLUnitTestProject
{
  [TestClass]
  public class TSFLUnitTest
  {
    [TestMethod][TestCategory("MA")]
    public void SMA_1D_Test()
    {
      //Send a small array of same length as period
      var array = new double[] {1.0, 2.0, 3.0};
      
      var result = TSFL.SMA.Compute(array, 3);
      //result should come back with a 2 in the last position (1+2+3)/3
      Assert.AreEqual(2.0, result[2]);

      //Send a small array of same length as period
      array = new double[] { 1.0, 2.0, 3.0, 4.0, 5.0 };
      result = TSFL.SMA.Compute(array, 3);
      //result should come back with a 4 in the last position (3+4+5)/3
      Assert.AreEqual(4.0, result[4]);

      //Call with additional start and end parameters set to beginning and end of array
      result = TSFL.SMA.Compute(array, 3, 0, array.Length);
      Assert.AreEqual(4.0, result[4]);

      //Call with intermediate range to compute, assert on 2nd to last
      result = TSFL.SMA.Compute(array, 3, 1, 3);
      Assert.AreEqual(3.0, result[3]);
    }
    [TestMethod][TestCategory("MA")]
    public void SMA_2D_Test()
    {
      //Send a small array of same length as period
      var array = new double[,] { {1.0, 1.0}, {2.0, 2.0}, {3.0, 3.0} };
      var result = TSFL.SMA.Compute(array, 3, 1);
      //result should come back with a 2 in the last position (1+2+3)/3
      Assert.AreEqual(2.0, result[2]);

      //Send a small array of same length as period
      array = new double[,] { { 1.0, 1.0 }, { 2.0, 2.0 }, { 3.0, 3.0 }, { 4.0, 4.0 }, { 5.0, 5.0 } };
      result = TSFL.SMA.Compute(array, 3, 1);
      //result should come back with a 4 in the last position (3+4+5)/3
      Assert.AreEqual(4.0, result[4]);

      //Call with additional start and end parameters set to beginning and end of array
      result = TSFL.SMA.Compute(array, 3, 0, array.Length, 1);
      Assert.AreEqual(4.0, result[4]);

      //Call with intermediate range to compute, assert on 2nd to last
      result = TSFL.SMA.Compute(array, 3, 1, 3, 1);
      Assert.AreEqual(3.0, result[3]);
    }
  }
}
