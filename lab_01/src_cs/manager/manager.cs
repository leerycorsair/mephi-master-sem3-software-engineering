using System;
using System.Collections.Generic;
using System.Linq;
using ResourceRegistrator.Model;

namespace ResourceRegistrator.Manager
{
    public class Manager
    {
        public RequestSlice Requests { get; set; }
        public Model.Model Model { get; set; }

        public Manager(Model.Model model)
        {
            Requests = new RequestSlice();
            Model = model;
        }

        public bool CheckModel()
        {
            return Model.IsEmpty();
        }

        public void SetModel(Model.Model model)
        {
            Model = model;
        }

        public void AddRequest(Request request)
        {
            var resource = Model.SearchResource(request.ResourceId);
            if (resource == null)
            {
                throw new Exception(Errors.ResourceNotFound);
            }
            Requests.Add(request);
        }

        public void FreeResource(string resourceId)
        {
            Model.FreeResource(resourceId);
            var req = Requests.FirstOrDefault(r => r.ResourceId == resourceId && r.Status == "Active");
            if (req != null)
            {
                Requests.Remove(req);
            }
        }

        public void Work()
        {
            foreach (var r in Requests)
            {
                if (r.Status == "Waiting")
                {
                    try
                    {
                        Model.OccupyResource(r.ResourceId);
                    }
                    catch (Exception ex)
                    {
                        if (ex.Message == Errors.ResourceIsBusy)
                        {
                            continue;
                        }
                    }
                    r.Status = "Active";
                }
                r.Time -= 1;
                if (r.Time <= 0)
                {
                    Model.FreeResource(r.ResourceId);
                }
            }
            Requests.RemoveAll(r => r.Time <= 0);
        }

        public void InitModel(int resourceCount)
        {
            Model.InitResources(resourceCount);
        }

        public void InitFromFile(string path)
        {
            Model.InitFromFile(path);
        }

        public string RequestsInfo()
        {
            return Requests.ToString();
        }

        public string ModelInfo()
        {
            return Model.ToString();
        }
    }
}
